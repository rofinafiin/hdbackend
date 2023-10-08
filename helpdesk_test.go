package hdbackend

import (
	"fmt"
	"github.com/aiteung/atdb"
	"os"
	"testing"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "HelpdeskData",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

//func TestInsertData(t *testing.T) {
//	sistem := "ITeung"
//	status := "Aktif"
//	bio := User{
//		Nama:      "Nafiis",
//		Email:     "rofinafiisr@gmail.com",
//		Handphone: "6285156007137",
//	}
//	hasil := InsertDataComp(MongoConn, sistem, status, bio)
//	fmt.Println(hasil)
//
//}
//
//func TestGetDataCompFromStatus(t *testing.T) {
//	stats := "Aktif"
//	biodata := GetDataCompFromStatus(stats, MongoConn, "data_complain")
//	fmt.Println(biodata)
//}
//
//func TestGetDataHelperFromPhone(t *testing.T) {
//	hp := "085156007137"
//	hasil := GetDataHelperFromPhone(hp, MongoConn, "helperdata")
//	fmt.Println(hasil)
//}
//
//func TestGetDataCompFromHandphone(t *testing.T) {
//	hp := "085156007137"
//	hasil := GetDataCompFromHandphone(hp, MongoConn, "data_complain")
//	fmt.Println(hasil)
//}
//
//func TestGetDataAllbyStats(t *testing.T) {
//	stats := "Aktif"
//	data := GetDataAllbyStats(stats, MongoConn, "data_complain")
//	fmt.Println(data)
//}
//
//func TestInsertDataHelper(t *testing.T) {
//	helpid := "115"
//	username := "rofigansbanget"
//	nama := "Rofi Nafiis"
//	email := "rofinafiisr@ulbi.ac.id"
//	handphone := "085156007137"
//	hasil := InsertDataHelper(helpid, username, nama, email, handphone, MongoConn, "helperdata")
//	fmt.Println(hasil)
//
//}
//
//func TestDeleteDataHelper(t *testing.T) {
//	hp := "085156007137"
//	res := DeleteDataHelper(hp, MongoConn, "helperdata")
//	fmt.Println(res)
//}

//func TestInsertJumlahComplain(t *testing.T) {
//
//	result := InsertJumlahComplain(MongoConn, )
//	fmt.Println(result)
//}

//func TestGetDataJumlah(t *testing.T) {
//	tahun := "2023"
//	hasil := GetDataJumlah(tahun, MongoConn, "jumlah_complain")
//	fmt.Println(hasil)
//}

func TestInsertCritics(t *testing.T) {
	name := "ropigantengbanget"
	value := "bagus banget"
	hasil := InsertCritics(MongoConn, name, value)
	fmt.Println(hasil)
}
