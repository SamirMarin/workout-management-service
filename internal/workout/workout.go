package workout

import (
	"github.com/SamirMarin/workout-management-service/internal/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	awsDynamoDb "github.com/aws/aws-sdk-go/service/dynamodb"
	awsDynamoDbAttribute "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strconv"
)

type Workout struct {
	Owner     string     `json:"owner"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	Equipment Equipment  `json:"equipment"`
	Exercises []Exercise `json:"exercises"`
}

type Equipment struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Exercise struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Sets        int    `json:"reps"`
	Time        int    `json:"time"`
}

// CreateWorkout creates a workout, save new workout to db
func (w *Workout) CreateWorkout() error {
	dynamoDbClient := dynamodb.NewClient("Workout")
	err := dynamoDbClient.StoreItem(w)
	if err != nil {
		return err
	}
	return nil
}
func (w *Workout) GetWorkout() error {
	dynamoDbClient := dynamodb.NewClient("Workout")
	err, getItemOutput := dynamoDbClient.GetItem(w)
	if err != nil {
		return err
	}
	err = awsDynamoDbAttribute.UnmarshalMap(getItemOutput.Item, w)
	return nil
}

func (w *Workout) ToDynamoDbAttribute() map[string]*awsDynamoDb.AttributeValue {
	exerciseList := make([]*awsDynamoDb.AttributeValue, len(w.Exercises))
	for i, exercise := range w.Exercises {
		exerciseList[i] = &awsDynamoDb.AttributeValue{
			M: map[string]*awsDynamoDb.AttributeValue{
				"Name": {
					S: aws.String(exercise.Name),
				},
				"Description": {
					S: aws.String(exercise.Description),
				},
				"Sets": {
					N: aws.String(strconv.Itoa(exercise.Sets)),
				},
				"Time": {
					N: aws.String(strconv.Itoa(exercise.Time)),
				},
			},
		}
	}
	return map[string]*awsDynamoDb.AttributeValue{
		"Owner": {
			S: aws.String(w.Owner),
		},
		"Name": {
			S: aws.String(w.Name),
		},
		"Category": {
			S: aws.String(w.Category),
		},
		"Equipment": {
			M: map[string]*awsDynamoDb.AttributeValue{
				"name": {
					S: aws.String(w.Equipment.Name),
				},
				"description": {
					S: aws.String(w.Equipment.Description),
				},
			},
		},
		"Exercises": {
			L: exerciseList,
		},
	}
}

func (w *Workout) ToDynamoDbItemInput() *awsDynamoDb.GetItemInput {
	return &awsDynamoDb.GetItemInput{
		TableName: aws.String("Workout"),
		Key: map[string]*awsDynamoDb.AttributeValue{
			"Owner": {
				S: aws.String(w.Owner),
			},
			"Name": {
				S: aws.String(w.Name),
			},
		},
	}
}
