#!/bin/bash

# Command to create a DynamoDB table
# Specifying ReadCapacityUnits and WriteCapacityUnits is required in local mode
aws dynamodb create-table \
    --table-name Workout \
    --attribute-definitions \
        AttributeName=Owner,AttributeType=S \
        AttributeName=Name,AttributeType=S \
    --key-schema \
        AttributeName=Owner,KeyType=HASH \
        AttributeName=Name,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --region us-west-2 \
    --endpoint-url http://localhost:8000