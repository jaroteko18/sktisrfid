package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	db "sktisrfid/pkg/database"
	"time"
)

type Person struct {
	RFIDID string
}

type PayloadInsertDelete struct {
	insert []PayloadAbsenteeism
	delete []PayloadAbsenteeism
}

type PayloadAbsenteeism struct {
	AbsentDate  string
	AbsentType  string
	RFIDID      string
	EmployeeID  string
	CreatedDate string
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

type ResponseResult struct {
	Status  string
	Message string
}

func ListAbsent(data map[string]interface{}) (res []ListAbsenteeism) {
	var payload PayloadAbsenteeism
	payload.AbsentDate = data["AbsentDate"].(string)
	payload.AbsentType = data["AbsentType"].(string)

	rows, err := db.DB.Query("select RFIDID,EP.EmployeeID,EmployeeNumber,EmployeeName, "+
		"Plant as LocationCode,[Group] as GroupCode,Unit as UnitCode,EP.CreatedDate "+
		"from [SKTIS].[dbo].[ExePlantWorkerAbsenteeism] EP "+
		"inner join [SKTIS].[dbo].[MstRFID] MR "+
		"on EP.EmployeeID=MR.EmployeeID "+
		"Where IsActive=1 and IsFromRFID=1 "+
		"and StartDateAbsent>=$2 and EndDateAbsent<=$2 ORDER BY EP.CreatedDate DESC", payload.AbsentType, payload.AbsentDate)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}
	for rows.Next() {
		list := ListAbsenteeism{}
		err = rows.Scan(&list.RFIDID, &list.EmployeeID, &list.EmployeeNumber, &list.EmployeeName, &list.LocationCode, &list.GroupCode, &list.UnitCode, &list.CreatedDate)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
			return
		}
		res = append(res, list)
	}
	err = rows.Err()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}
	return

}

func InsertDeleteAbsent(data map[string]interface{}) (res ResponseResult) {
	var payload PayloadInsertDelete

	ParamInsert, _ := json.Marshal(data["insert"])
	if err := json.Unmarshal(ParamInsert, &payload.insert); err != nil {
		log.Fatal(err)
		res.Status = "error"
		res.Message = err.Error()
		return
	}

	ParamDelete, _ := json.Marshal(data["delete"])
	if err := json.Unmarshal(ParamDelete, &payload.delete); err != nil {
		log.Fatal(err)
		res.Status = "error"
		res.Message = err.Error()
		return
	}

	if payload.delete != nil {
		for i := 0; i < len(payload.delete); i++ {

			_, err := db.DB.Query("EXEC [SKTIS].[dbo].[DELETE_WORKER_ABSENTEEISM_RFID] @AbsentDate = $1 ,@AbsentType = $2, @RFIDID = $3, @EmployeeID = $4, @CreatedDate = $5, @CreatedBy = $6, @UpdatedBy = $7, @UpdatedDate = $8", payload.delete[i].AbsentDate, payload.delete[i].AbsentType, payload.delete[i].RFIDID, payload.delete[i].EmployeeID, payload.delete[i].CreatedDate, "RFID", "RFID", payload.delete[i].CreatedDate)

			if err != nil {
				log.Fatal(err)
				res.Status = "error"
				res.Message = err.Error()
				return

			}

		}
	}

	if payload.insert != nil {
		for i := 0; i < len(payload.insert); i++ {
			_, err := db.DB.Query("EXEC [SKTIS].[dbo].[INSERT_WORKER_ABSENTEEISM_RFID] @AbsentDate = $1 ,@AbsentType = $2, @RFIDID = $3, @EmployeeID = $4, @CreatedDate = $5, @CreatedBy = $6, @UpdatedBy = $7, @UpdatedDate = $8", payload.insert[i].AbsentDate, payload.insert[i].AbsentType, payload.insert[i].RFIDID, payload.insert[i].EmployeeID, payload.insert[i].CreatedDate, "RFID", "RFID", payload.insert[i].CreatedDate)

			if err != nil {
				log.Fatal(err)
				res.Status = "error"
				res.Message = err.Error()
				return
			}
		}
	}

	res.Status = "success"
	res.Message = "Data was successfully saved !"

	return
}
