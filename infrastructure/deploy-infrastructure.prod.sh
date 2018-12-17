#!/bin/bash

export AWS_PROFILE=prod
aws cloudformation \
  deploy \
  --template-file ac-sqs-lambda.cf-template.json \
  --stack-name ac-sqs-lambda-prod \
  --parameter-overrides EnvironmentParameter=production \
  --capabilities CAPABILITY_NAMED_IAM
