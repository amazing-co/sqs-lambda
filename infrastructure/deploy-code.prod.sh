#!/bin/bash

aws s3 \
  cp \
  ac-sqs-lambda.zip \
  s3://ac-sqs-lambda-prod \
  --profile prod
