package controllers

import (
	"fmt"
	"log"
	db "sktisrfid/pkg/database"
	"time"
)

type Person struct {
	RFIDID string
}

type PayloadAbsenteeism struct {
	AbsentDate string
	AbsentType string
}

type ListAbsenteeism struct {
	RFIDID         string
	EmployeeID     string
	EmployeeNumber string
	EmployeeName   string
	LocationCode   string
	GroupCode      string
	UnitCode       string
	CreatedDate    time.Time
}

func ListAbsent(data map[string]interface{}) (res []ListAbsenteeism) {
	var payload PayloadAbsenteeism
	payload.AbsentDate = data["AbsentDate"].(string)
	payload.AbsentType = data["AbsentType"].(string)

	rows, err := db.DB.Query("select RFIDID,EP.EmployeeID,EmployeeNumber,EmployeeName, "+
		"Plant as LocationCode,[Group] as GroupCode,Unit as UnitCode,EP.CreatedDate "+
		"from ExePlantWorkerAbsenteeism EP "+
		"inner join MstRFID MR "+
		"on EP.EmployeeID=MR.EmployeeID "+
		"Where IsActive=1 and IsFromRFID=1 "+
		"and AbsentType=$1 and StartDateAbsent>=$2 and EndDateAbsent<=$2;", payload.AbsentType, payload.AbsentDate)

	for rows.Next() {
		list := ListAbsenteeism{}
		err = rows.Scan(&list.RFIDID, &list.EmployeeID, &list.EmployeeNumber, &list.EmployeeName, &list.LocationCode, &list.GroupCode, &list.UnitCode, &list.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(list)

		res = append(res, list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return

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
