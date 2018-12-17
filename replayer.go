package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type replayer struct {
	session    *session.Session
	queueName  string
	lambdaName string
}

func (r replayer) process() error {
	fmt.Println("Replaying messages...")

	// Create services client.
	sqsSvc := sqs.New(r.session)
	lambdaSvc := lambda.New(r.session)

	// Need to convert the queue name into a URL. Make the GetQueueUrl
	// API call to retrieve the URL. This is needed for receiving messages
	// from the queue.
	resultURL, err := sqsSvc.GetQueueUrl(
		&sqs.GetQueueUrlInput{
			QueueName: aws.String(r.queueName),
		},
	)
	if err != nil {
		return fmt.Errorf("Error getting the queue: %s details, error: %+v", r.queueName, err)
	}

	// Receive a message from the SQS queue with long polling enabled.
	sqsRcvInput := &sqs.ReceiveMessageInput{
		QueueUrl:            resultURL.QueueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	}

	lambdaFunc := func(body []byte) {
		li := &lambda.InvokeInput{
			ClientContext:  aws.String("Replayer Processor"),
			FunctionName:   aws.String(r.lambdaName),
			InvocationType: aws.String("Event"),
			LogType:        aws.String("Tail"),
			Payload:        []byte(body),
		}

		_, err = lambdaSvc.Invoke(li)
		if err != nil {
			// Log & Skip message when error
			fmt.Printf("Unable to process lambda, Error: %+v", err)
		}

		// fmt.Println(lambdaOut)
	}

	doneFunc := func(receiptHandle *string) {
		_, err := sqsSvc.DeleteMessage(
			&sqs.DeleteMessageInput{
				QueueUrl:      resultURL.QueueUrl,
				ReceiptHandle: receiptHandle,
			})
		if err != nil {
			// Skip message when nil
			fmt.Printf("Unable to delete message from queue: %s, Error: %+v \n", *resultURL.QueueUrl, err)
		}

		// fmt.Println(sqsDelOut)
	}

	for {
		done := doWork(sqsSvc, sqsRcvInput, lambdaFunc, doneFunc)
		if !done {
			break
		}
	}

	return nil
}

func doWork(
	sqsSvc *sqs.SQS,
	sqsRcvInput *sqs.ReceiveMessageInput,
	lambdaFunc func(body []byte),
	doneFunc func(receiptHandle *string),
) bool {

	sqsRcvOut, err := sqsSvc.ReceiveMessage(sqsRcvInput)
	if err != nil {
		fmt.Printf("Unable to receive message from %+v, Error: %+v", sqsRcvInput, err)
		return false
	}

	if len(sqsRcvOut.Messages) < 1 {
		return false
	}

	for _, m := range sqsRcvOut.Messages {
		receiptHandle := m.ReceiptHandle
		bodyBytes := []byte(*m.Body)

		fmt.Printf("Processing: %s \n", *receiptHandle)
		fmt.Printf("Body: %s \n", *m.Body)

		// Skip message when nil
		if receiptHandle == nil || bodyBytes == nil {
			fmt.Printf("Error Receipt or Body is nil")
			continue
		}

		// Run the largeted ambda
		lambdaFunc(bodyBytes)

		// Delete when done
		doneFunc(receiptHandle)
	}
	return true
}
