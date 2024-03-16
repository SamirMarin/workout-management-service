#!/bin/bash
echo "Waiting for DynamoDB Local to be ready..."
# Wait for DynamoDB Local to become ready
until aws dynamodb list-tables --endpoint-url http://localhost:8000 --region us-west-2 > /dev/null 2>&1; do
    echo "DynamoDB Local is not ready yet..."
    sleep 5
done
echo "DynamoDB Local is ready."
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