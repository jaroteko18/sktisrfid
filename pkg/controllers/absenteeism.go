package controllers

import (
	"fmt"
	db "sktisrfid/pkg/database"
)

type Person struct {
	RFIDID string
}

type Test struct {
}

func GetListAbsenteeism(empid string) string {
	row := db.DB.QueryRow("SELECT RFIDID FROM [MstRFID] WHERE EmployeeID=$1;", empid)
	fmt.Println(row)
	p := new(Person)
	err := row.Scan(&p.RFIDID)
	if err != nil {
		fmt.Println("empid", empid)
		fmt.Println("err=>", err)
		return "Gagal dapat data"
	} else {
		fmt.Println(p)
		return p.RFIDID
	}
}

// func TestRFID(list map[string]interface{}) string {
// 	st := []Test{}
// 	row := db.DB.Selec .QueryRow(&st, "SELECT RFIDID, EmployeeID FROM [MstRFID]")
// 	fmt.Println(row)
// 	p := new(Person)
// 	err := row.Scan(&p.RFIDID)
// 	if err != nil {
// 		fmt.Println("empid", empid)
// 		fmt.Println("err=>", err)
// 		return "Gagal dapat data"
// 	} else {
// 		fmt.Println(p)
// 		return p.RFIDID
// 	}
// }
