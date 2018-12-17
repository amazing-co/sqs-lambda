# sqs-lambda

A lambda function that replays messages from an SQS(Simple Query Service) queue to another lambda function

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

- [Go](https://golang.org/)

### Installing

Compile the code

```go
go build
```

an example of getting some data out of the system or using it for a little demo - TBC

## Running the tests

Unit test with [Go - Test packages](https://golang.org/cmd/go/#hdr-Test_packages)

run unit tests

```go
go test ./...
```

### And coding style tests

- Code Style: [Effective Go](https://golang.org/doc/effective_go.html)
- Linter: [Golint](https://github.com/golang/lint)

## Deployment

1.Compress the code to `zip` file and upload it to AWS Lambda or S3 Bucket

```bash
cd infrastructure
./compress.sh
./deploy-code.prod.sh
```

3.Create your own CloudFormation stack `ac-sqs-lambda.cf-template.json`

4.Deploy Your Own Infrastructure

```bash
cd infrastructure
./deploy-infrastructure.prod.sh
```

Alternatively, You can just create your own lambda infrastructure on AWS nad and
upload [`ac-sqs-lambda.zip`](https://github.com/amazing-co/sqs-lambda/blob/master/infrastructure/ac-sqs-lambda).

NOTE!
You need to setup 3 different variables:

- REGION
- SQS_QUEUE_SOURCE_NAME
- LAMBDA_TARGET_NAME

## Built With

- [AWS Lambda](https://aws.amazon.com/lambda/)
- [AWS CloudFormation](https://docs.aws.amazon.com/cloudformation/)
- [AWS S3](https://aws.amazon.com/s3/)

## Contributing

TBC

## Versioning

TBC

## Authors

- **Raymond Boles** - _Initial work_ - [AmazingCo](https://github.com/amazing-co)

See also the list of [contributors](https://github.com/amazing-co/ac-sqs-lambda/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License

## Acknowledgments

TBC
