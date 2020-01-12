package weinvite

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Table interface {
	Name() string          // for name table
	Structur() interface{} //  for declaring struct
}

func Connect() (*mongo.Database, error) {
	ctx := context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	return client.Database("weinvite"), nil
}

func Insert(ctx context.Context, t Table) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	fmt.Println(t.Structur())

	_, err = db.Collection(t.Name()).InsertOne(ctx, t.Structur())
	if err != nil {
		return err
	}

	return nil
}

//still developing
func Find(ctx context.Context, filter bson.M, t Table) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	csr, err := db.Collection(t.Name()).Find(ctx, filter)
	if err != nil {
		return err
	}
	// defer csr.Close(ctx)
	// result := make(map[string]interface{}, 0)
	for csr.Next(ctx) {
		row := make(map[string]interface{}, 0)
		if err := csr.Decode(&row); err != nil {
			fmt.Println("ini massuk")
			return err
		}
		// result := make(row)

		fmt.Println(row)
		// result = row
		// fmt.Println(result)
	}

	fmt.Println("pass")
	return nil

}

func Update(ctx context.Context, change map[string]interface{}, selector bson.M, t Table) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	_, err = db.Collection(t.Name()).UpdateOne(ctx, selector, bson.M{"$set": change})
	if err != nil {
		return err
	}
	return nil
}

func Remove(ctx context.Context, t Table, selector bson.M) error {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection(t.Name()).DeleteOne(ctx, selector)
	if err != nil {
		return err
	}
	return nil
}
