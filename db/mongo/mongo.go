package mongo

import (
	"context"
	"micro_services/msbase/logger"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (rcv DbClient) Create(collection string, model interface{}) error {
	act, err := rcv.Client.Database(Database).Collection(collection).InsertOne(context.TODO(), model)
	if err != nil {
		logger.Error("cannot insert data to the db ")
		return err
	}
	stringId := act.InsertedID.(primitive.ObjectID).Hex()
	logger.Infof("document inserted with id %v", stringId)
	return err
}

func (rcv DbClient) Read(collection string, q bson.M, result interface{}, params ...interface{}) error {
	var (
		err         error
		sort        string
		limit, page int64
	)

	opt := new(options.FindOptions)
	if len(params) > 0 {
		limit = int64(params[0].(int))
		if limit > 0 {
			opt.Limit = &limit
		}
	}

	if len(params) > 1 {
		page = int64(params[1].(int))
		if page > 0 {
			skip := (page - 1) * limit
			opt.Skip = &skip
		}
	}

	if len(params) > 2 {
		sort = params[2].(string)

		if sort != "" {
			splits := strings.Split(sort, ",")
			sortOpts := bson.D{}

			for _, s := range splits {
				if strings.HasPrefix(sort, "-") {
					sort = strings.ReplaceAll(s, "-", "")
					// sortOpts = append(sortOpts, bson.D{Name: s, Value: -1})
				} else {
					// sortOpts = append(sortOpts, bson.DocElem{Name: s, Value: 1})
				}
			}
			opt.SetSort(sortOpts)
		}

	}

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	_ = cancel
	cursor, err := rcv.Client.Database(Database).Collection(collection).Find(ctx, q, opt)
	if err != nil {
		logger.Error("error : ", err.Error())
		return err
	}
	defer cursor.Close(ctx)
	err = cursor.All(context.Background(), result)
	if err != nil {
		logger.Error("error : ", err.Error())
	}
	return nil
}

func (rcv DbClient) ReadOne(collection string, q bson.M, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	_ = cancel
	return rcv.Client.Database(Database).Collection(collection).FindOne(ctx, q).Decode(result)
}

func (rcv DbClient) UpdateOne(collection string, filter bson.M, update bson.M) error {
	var (
		err error
	)
	if _, err = rcv.Client.Database(Database).Collection(collection).UpdateOne(context.TODO(), filter, update); err != nil {
		logger.Error("cannot update data to the db, error : ", err.Error())
	}

	return err
}

func (rcv DbClient) UpdateMany(collection string, filter bson.M, update bson.M) error {
	var (
		err error
	)
	if _, err = rcv.Client.Database(Database).Collection(collection).UpdateMany(context.TODO(), filter, update); err != nil {
		logger.Error("cannot update many data to the db, error : ", err.Error())
	}

	return err
}

func (rcv DbClient) DeleteOne(collection string, filter bson.M) error {
	_, err := rcv.Client.Database(Database).Collection(collection).DeleteOne(context.TODO(), filter)
	return err
}

func (rcv DbClient) DeleteMany(collection string, filter bson.M) error {
	_, err := rcv.Client.Database(Database).Collection(collection).DeleteMany(context.TODO(), filter)
	return err
}

func (rcv DbClient) Count(collection string, filter bson.M) (int, error) {
	count, err := rcv.Client.Database(Database).Collection(collection).CountDocuments(context.TODO(), filter)
	return int(count), err
}
