package main

import (
	"fmt"
	"os"

	lambdaGo "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type processor interface {
	process() error
}

var processors []processor

func init() {
	var region = os.Getenv("REGION")
	var queueName = os.Getenv("SQS_QUEUE_SOURCE_NAME")
	var lambdaName = os.Getenv("LAMBDA_TARGET_NAME")

	// Initialise a session in a AWS region that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials or lambda role.
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
	panicErrorf(err)

	processors = []processor{
		replayer{
			session:    sess,
			queueName:  queueName,
			lambdaName: lambdaName,
		},
	}
	panicErrorf(err)
}

func handler() (string, error) {
	for _, p := range processors {
		panicErrorf(p.process())
	}
	return "Finished!", nil
}

func main() {
	lambdaGo.Start(handler)
}

func panicErrorf(err error) {
	if err != nil {
		fmt.Printf("Error: %+v \n", err)
		panic(err)
	}
}
