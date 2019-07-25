package account

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	COLLECTION = "accounts"
)

func (db *repository) create(i interface{}) error {
	collection := db.mongo.Database(db.database).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), i)
	if err != nil {
		return err
	}
	return nil
}
func (db *repository) read(key string, value string, i interface{}) error {
	collection := db.mongo.Database(db.database).Collection(COLLECTION)
	filter := bson.D{{key, value}}
	err := collection.FindOne(context.TODO(), filter).Decode(&i)
	if err != nil {
		return err
	}
	return nil
}
func (db *repository) update(key string, value string, field interface{}) (interface{}, error) {
	collection := db.mongo.Database(db.database).Collection(COLLECTION)
	filter := bson.D{{key, value}}
	toUpdate := bson.D{
		{"$set", bson.D{
			{"balance", field},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, toUpdate, nil)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}
func (db *repository) delete(key string, value string, i interface{}) (interface{}, error) {
	collection := db.mongo.Database(db.database).Collection(COLLECTION)
	filter := bson.D{{key, value}}
	deleted, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}
func (db *repository) readAll() ([]*Account, error) {
	accounts := make([]*Account, 0)
	collection := db.mongo.Database(db.database).Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem Account
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		accounts = append(accounts, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return accounts, nil
}
