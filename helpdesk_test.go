package HelpdeskData

import (
	"fmt"
	"testing"
)

//func TestInsertData(t *testing.T) {
//	sistem := "ITeung"
//	status := "Aktif"
//	bio := User{
//		Nama:      "Nafiis",
//		Email:     "rofinafiisr@gmail.com",
//		Handphone: "6285156007137",
//	}
//	hasil := InsertDataComp(sistem, status, bio)
//	fmt.Println(hasil)
//
//}
//
//func TestGetKaryawanFromStatus(t *testing.T) {
//	stats := "Aktif"
//	biodata := GetDataCompFromStatus(stats)
//	fmt.Println(biodata)
//}

//func TestGetDataUserFromPhone(t *testing.T) {
//	phone := "6285156007137"
//	biodata := GetDataUserFromPhone(phone)
//	fmt.Println(biodata)
//}

func TestGetDataAllbyStats(t *testing.T) {
	stats := "Aktif"
	data := GetDataAllbyStats(stats)
	fmt.Println(data)
}
