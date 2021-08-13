package controllers

import (
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
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		list := ListAbsenteeism{}
		err = rows.Scan(&list.RFIDID, &list.EmployeeID, &list.EmployeeNumber, &list.EmployeeName, &list.LocationCode, &list.GroupCode, &list.UnitCode, &list.CreatedDate)
		if err != nil {
			log.Fatal(err)
			return
		}
		res = append(res, list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return
	}
	return

}
