package mongo

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoCli *mongo.Client
	DBClient *DbClient
)

var (
	Database string
)

type DbClient struct {
	Client *mongo.Client
	Retry  string
}
