package backendjulgo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func GadgetInsertOne(gadget Gadget) (interface{}, error) {
	insertResult, err := MongoConnect("dzul2024").Collection("gadget").InsertOne(context.Background(), gadget)
	if err != nil {
		return nil, fmt.Errorf("error inserting gadget: %v", err)
	}
	return insertResult.InsertedID, nil
}


func GadgetGetAll() ([]Gadget, error) {
	var gadgets []Gadget

	cursor, err := MongoConnect("dzul2024").Collection("gadget").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching gadgets: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var gadget Gadget
		err := cursor.Decode(&gadget)
		if err != nil {
			return nil, fmt.Errorf("error decoding gadget: %v", err)
		}
		gadgets = append(gadgets, gadget)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return gadgets, nil
}


func PenggunaInsertOne(pengguna Pengguna) (interface{}, error) {
	insertResult, err := MongoConnect("dzul2024").Collection("Pengguna").InsertOne(context.Background(), pengguna)
	if err != nil {
		return nil, fmt.Errorf("error inserting pengguna: %v", err)
	}
	return insertResult.InsertedID, nil
}


func PenggunaGetAll() ([]Pengguna, error) {
	var penggunas []Pengguna

	cursor, err := MongoConnect("dzul2024").Collection("Pengguna").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching penggunas: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var pengguna Pengguna
		err := cursor.Decode(&pengguna)
		if err != nil {
			return nil, fmt.Errorf("error decoding pengguna: %v", err)
		}
		penggunas = append(penggunas, pengguna)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return penggunas, nil
}


func TransaksiInsertOne(transaksi Transaksi) (interface{}, error) {
	insertResult, err := MongoConnect("dzul2024").Collection("Transaksi").InsertOne(context.Background(), transaksi)
	if err != nil {
		return nil, fmt.Errorf("error inserting transaksi: %v", err)
	}
	return insertResult.InsertedID, nil
}


func TransaksiGetAll() ([]Transaksi, error) {
	var transaksis []Transaksi

	cursor, err := MongoConnect("dzul2024").Collection("Transaksi").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching transaksis: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var transaksi Transaksi
		err := cursor.Decode(&transaksi)
		if err != nil {
			return nil, fmt.Errorf("error decoding transaksi: %v", err)
		}
		transaksis = append(transaksis, transaksi)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return transaksis, nil
}
