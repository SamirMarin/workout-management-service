package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

type Storable interface {
	ToDynamoDbAttribute() map[string]*dynamodb.AttributeValue
}

type Client struct {
	Dynamodb  *dynamodb.DynamoDB
	TableName string
}

func NewClient(tableName string) *Client {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Optional: Override with local endpoint if an environment variable is set
	var svc *dynamodb.DynamoDB
	if localEndpoint := os.Getenv("DYNAMODB_LOCAL_ENDPOINT"); localEndpoint != "" {
		svc = dynamodb.New(sess, &aws.Config{
			Endpoint: aws.String(localEndpoint),
			Region:   aws.String("us-west-2"),
			// provide dummy credentials when connecting to DynamoDB local
			Credentials: credentials.NewStaticCredentials("dummy", "dummy", ""),
			// Disable SSL for local non-production use
			DisableSSL: aws.Bool(true),
		})
	} else {
		svc = dynamodb.New(sess)
	}

	return &Client{
		Dynamodb:  svc,
		TableName: tableName,
	}
}

func (c *Client) StoreItem(itemToStore Storable) error {
	item := itemToStore.ToDynamoDbAttribute()

	_, err := c.Dynamodb.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(c.TableName),
	})
	if err != nil {
		return err
	}

	return nil
}
