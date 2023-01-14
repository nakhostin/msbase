package mongo

import (
	"context"
	"micro_services/msbase/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) {

	clientInstall, err := Client(uri)
	if err != nil {
		panic(err)
	}
	MongoCli = clientInstall
	DBClient = &DbClient{
		Client: clientInstall,
	}
	Ticker := time.NewTicker(time.Duration(2) * time.Second)
	go func(client *mongo.Client) {
		for t := range Ticker.C {
			_ = t
			if err = client.Ping(context.TODO(), nil); err != nil {
				logger.Errorf("can not ping to the db with URI %s", "")
				logger.Info("Try to Connect DB")
				clientInstall, err := Client(uri)
				if err != nil {
					logger.Error("can not connect to MongoDB with error : ", err.Error())
				}
				MongoCli = clientInstall
				DBClient = &DbClient{
					Client: clientInstall,
				}
			}
		}
	}(MongoCli)
}

//Client will Connect to the mongo db and pass the client
func Client(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {

		return nil, err
	}

	DbAttentive(*client)
	return client, nil
}

//DbAttentive will check database and collection existence .If they are not exist in DB it will create them
func DbAttentive(client mongo.Client) error {
	if _, err := client.ListDatabaseNames(context.TODO(), bson.M{"name": Database}); err != nil {
		panic(err)
	}

	return nil
}
