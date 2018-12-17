#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o ac-sqs-lambda ../*.go
zip ac-sqs-lambda.zip ac-sqs-lambda
