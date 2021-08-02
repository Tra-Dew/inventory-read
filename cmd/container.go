package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/d-leme/tradew-inventory-read/pkg/core"
	"github.com/d-leme/tradew-inventory-read/pkg/inventory"
	"github.com/d-leme/tradew-inventory-read/pkg/inventory/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Container ...
type Container struct {
	Settings *core.Settings

	Authenticate *core.Authenticate

	MongoClient *mongo.Client

	SNS *session.Session
	SQS *session.Session

	InventoryRepository inventory.Repository
	InventoryController inventory.Controller
}

// NewContainer creates new instace of Container
func NewContainer(settings *core.Settings) *Container {

	container := new(Container)

	container.Settings = settings

	container.SQS = core.NewSession(
		settings.SQS.Region,
		settings.SQS.Endpoint,
		settings.SQS.Path,
		settings.SQS.Profile,
		settings.SQS.Fake,
	)

	container.SNS = core.NewSession(
		settings.SNS.Region,
		settings.SNS.Endpoint,
		settings.SNS.Path,
		settings.SNS.Profile,
		settings.SNS.Fake,
	)

	container.MongoClient = connectMongoDB(settings.MongoDB)

	container.Authenticate = core.NewAuthenticate(settings.JWT.Secret)

	container.InventoryRepository = mongodb.NewRepository(container.MongoClient, settings.MongoDB.Database)
	container.InventoryController = inventory.NewController(container.Authenticate, container.InventoryRepository)

	return container
}

// Controllers maps all routes and exposes them
func (c *Container) Controllers() []core.Controller {
	return []core.Controller{
		&c.InventoryController,
	}
}

// Close terminates every opened resources
func (c *Container) Close() {
	c.MongoClient.Disconnect(context.Background())
}

func connectMongoDB(conf *core.MongoDBConfig) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.ConnectionString))

	if err != nil {
		logrus.
			WithError(err).
			Fatal("error connecting to MongoDB")
	}

	client.Connect(context.Background())

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		logrus.
			WithError(err).
			Fatal("error pinging MongoDB")
	}

	return client
}
