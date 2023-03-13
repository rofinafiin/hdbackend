package hdbackend

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataComp(db *mongo.Database, sistem string, status string, bio User) (InsertedID interface{}) {
	var datacomp DataComplain
	datacomp.Sistemcomp = sistem
	datacomp.Status = status
	datacomp.Biodata = bio
	return InsertOneDoc(db, "data_complain", datacomp)
}

func InsertDataHelper(helper string, username string, nama string, email string, handphone string, db *mongo.Database, col string) (InsertedID interface{}) {
	help := new(Helper)
	help.Helpid = helper
	help.Username = username
	help.Nama = nama
	help.Email = email
	help.Handphone = handphone
	return InsertOneDoc(db, col, help)
}

func GetDataCompFromStatus(status string, db *mongo.Database, col string) (data DataComplain) {
	user := db.Collection(col)
	filter := bson.M{"status": status}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataCompFromPhoneNumber: %v\n", err)
	}
	return data
}

func GetDataAllbyStats(stats string, db *mongo.Database, col string) (data []DataComplain) {
	user := db.Collection(col)
	filter := bson.M{"status": stats}
	err, _ := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	return data
}

func GetDataHelperFromPhone(phone string, db *mongo.Database, col string) (data Helper) {
	user := db.Collection(col)
	filter := bson.M{"handphone": phone}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataCompFromPhoneNumber: %v\n", err)
	}
	return data
}

func DeleteDataHelper(phone string, db *mongo.Database, col string) (data Helper) {
	user := db.Collection(col)
	filter := bson.M{"handphone": phone}
	err, _ := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataHelper : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
