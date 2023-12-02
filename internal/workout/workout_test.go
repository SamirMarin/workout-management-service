package workout

import (
	"github.com/aws/aws-sdk-go/aws"
	awsDynamoDb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToDynamoDbAttribute(t *testing.T) {
	workout := &Workout{
		Owner:    "testUser",
		Name:     "testWorkout",
		Category: "testCategory",
		Equipment: Equipment{
			Name:        "testEquipment",
			Description: "testDescription",
		},
		Exercises: []Exercise{
			{
				Name:        "testExercise",
				Description: "testDescription",
				Sets:        1,
				Time:        1,
			},
		},
	}
	dynamodbAttribute := workout.ToDynamoDbAttribute()
	expectedDynamodbAttribute := map[string]*awsDynamoDb.AttributeValue{
		"Owner": {
			S: &workout.Owner,
		},
		"Name": {
			S: &workout.Name,
		},
		"Category": {
			S: &workout.Category,
		},
		"Equipment": {
			M: map[string]*awsDynamoDb.AttributeValue{
				"Name": {
					S: &workout.Equipment.Name,
				},
				"Description": {
					S: &workout.Equipment.Description,
				},
			},
		},
		"Exercises": {
			L: []*awsDynamoDb.AttributeValue{
				{
					M: map[string]*awsDynamoDb.AttributeValue{
						"Name": {
							S: &workout.Exercises[0].Name,
						},
						"Description": {
							S: &workout.Exercises[0].Description,
						},
						"Sets": {
							N: aws.String("1"),
						},
						"Time": {
							N: aws.String("1"),
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expectedDynamodbAttribute, dynamodbAttribute)
}

func TestToDynamoDbItemInput(t *testing.T) {
	workout := Workout{
		Owner: "testUser",
		Name:  "testWorkout",
	}
	dynamodbItemInput := workout.ToDynamoDbItemInput()
	expectedDynamodbItemInput := &awsDynamoDb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*awsDynamoDb.AttributeValue{
			"Owner": {
				S: &workout.Owner,
			},
			"Name": {
				S: &workout.Name,
			},
		},
	}
	assert.Equal(t, expectedDynamodbItemInput, dynamodbItemInput)
}
