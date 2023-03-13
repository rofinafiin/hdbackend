package HelpdeskData

import (
	"fmt"
	"testing"
)

func TestInsertData(t *testing.T) {
	sistem := "ITeung"
	status := "Aktif"
	bio := User{
		Nama:      "Nafiis",
		Email:     "rofinafiisr@gmail.com",
		Handphone: "6285156007137",
	}
	hasil := InsertDataComp(sistem, status, bio)
	fmt.Println(hasil)

}

func TestGetDataCompFromStatus(t *testing.T) {
	stats := "Aktif"
	biodata := GetDataCompFromStatus(stats)
	fmt.Println(biodata)
}

func TestGetDataHelperFromPhone(t *testing.T) {
	hp := "085156007137"
	hasil := GetDataHelperFromPhone(hp)
	fmt.Println(hasil)
}

func TestGetDataAllbyStats(t *testing.T) {
	stats := "Aktif"
	data := GetDataAllbyStats(stats)
	fmt.Println(data)
}

func TestInsertDataHelper(t *testing.T) {
	helpid := "115"
	username := "rofigansbanget"
	nama := "Rofi Nafiis"
	email := "rofinafiisr@ulbi.ac.id"
	handphone := "085156007137"
	hasil := InsertDataHelper(helpid, username, nama, email, handphone)
	fmt.Println(hasil)

}

//func TestDeleteDataHelper(t *testing.T) {
//	hp := "085156007137"
//	res := DeleteDataHelper(hp)
//	fmt.Println(res)
//}
