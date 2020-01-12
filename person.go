package weinvite

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Nama   string `json:"username"bson:"username"`
	Alamat string `json:"alamat"bson:"alamat"`
	Status bool   `json:"status"bson:"status"`
}

func (p *Person) Name() string {
	return "person"
}

func (p *Person) Structur() interface{} {
	return &Person{Nama: p.Nama, Alamat: p.Alamat, Status: p.Status}
}

func (p *Person) Insert(ctx context.Context) error {
	return Insert(ctx, p)
}

//still developing for cenral library
// func (p *Person) Find(ctx context.Context, Filter bson.M) error {
// 	return Find(ctx, Filter, p)
// }

func (p *Person) Update(ctx context.Context, change map[string]interface{}, selector bson.M) (map[string]interface{}, error) {
	return change, Update(ctx, change, selector, p)
}

func (p *Person) Remove(ctx context.Context, selector bson.M) error {
	return Remove(ctx, p, selector)
}

func FindAll() ([]Person, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	csr, err := db.Collection("person").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer csr.Close(context.Background())

	res := make([]Person, 0)

	for csr.Next(context.Background()) {
		var row Person
		if err := csr.Decode(&row); err != nil {
			return nil, err
		}

		res = append(res, row)
	}
	return res, nil

}

//prototype
func FindFilter(filter bson.M) ([]Person, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	csr, err := db.Collection("person").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer csr.Close(context.Background())

	res := make([]Person, 0)

	for csr.Next(context.Background()) {
		var row Person
		if err := csr.Decode(&row); err != nil {
			return nil, err
		}

		res = append(res, row)
	}
	return res, nil

}
