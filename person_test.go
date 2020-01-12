package weinvite

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestInsert(t *testing.T) {
	var dummyData = []*Person{
		&Person{Nama: "filza", Alamat: "Kayuringin"},
		&Person{Nama: "Rahmi", Alamat: "Kayuringin"},
		&Person{Nama: "Fauziah", Alamat: "Kayuringin"},
	}
	t.Run("testing untuk insert", func(t *testing.T) {
		for _, item := range dummyData {
			ctx := context.Background()
			if err := item.Insert(ctx); err != nil {
				t.Fatalf("error:%v", err)
			}
			fmt.Println("Sukses")
		}
	})
}

func TestFindFilter(t *testing.T) {
	filter := bson.M{"status": false}

	res, err := FindFilter(filter)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}

// func TestFind(t *testing.T) {
// 	filter := bson.M{"username": "filza"}
// 	var temp Person
// 	ctx := context.Background()
// 	err := temp.Find(ctx, filter)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// }

func TestUpdate(t *testing.T) {
	selector := bson.M{"username": "filza"}
	change := []map[string]interface{}{
		{"status": true},
	}
	temp := &Person{}

	for _, item := range change {
		_, err := temp.Update(context.Background(), item, selector)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("sukses Update")
	}

}

func TestDelete(t *testing.T) {
	selector := bson.M{"username": "filza"}
	temp := &Person{}

	if err := temp.Remove(context.Background(), selector); err != nil {
		t.Fatal(err)
	}

	res, err := FindFilter(bson.M{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}

func TestFindAll(t *testing.T) {
	res, err := FindAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range res {
		fmt.Println(item)
	}
}
