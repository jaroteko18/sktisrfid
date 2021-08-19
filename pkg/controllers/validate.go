package controllers

import (
	"database/sql"
	"fmt"
	"log"
	db "sktisrfid/pkg/database"
	"time"
)

type PayloadValidate struct {
	Date       string
	AbsentType string
	RFIDID     string
}

type DetailRFID struct {
	RFIDID         string
	EmployeeID     string
	EmployeeNumber string
	EmployeeName   string
	LocationCode   string
	GroupCode      string
	UnitCode       string
	ProdCapacity   float32
	ProdTarget     float32
	CreatedDate    time.Time
}

type MasterRFID struct {
	RFIDID         string
	EmployeeID     string
	EmployeeName   string
	EmployeeNumber string
	LocationCode   string
	GroupCode      string
	UnitCode       string
}

type ResponseValidate struct {
	Message string
	Data    DetailRFID
	Status  string
}

func ValidateItem(data map[string]interface{}) (res ResponseValidate) {
	var payload PayloadValidate
	payload.Date = data["Date"].(string)
	payload.AbsentType = data["AbsentType"].(string)
	payload.RFIDID = data["RFIDID"].(string)

	listMst := MasterRFID{}
	err := db.DB.QueryRow("SELECT RFIDID, EmployeeID, NoPengenal as EmployeeNumber, EmployeeName, "+
		"Plant as LocationCode, Unit as UnitCode, [Group] as GroupCode "+
		"FROM [SKTIS].[dbo].[MstRFID] WHERE RFIDID=$1", payload.RFIDID).Scan(&listMst.RFIDID, &listMst.EmployeeID, &listMst.EmployeeNumber,
		&listMst.EmployeeName, &listMst.LocationCode, &listMst.UnitCode, &listMst.GroupCode)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		res.Message = "Data not found !"
		res.Status = "error"
		res.Data = DetailRFID{}
		fmt.Println("MASOK 1")
		return
	}
	fmt.Println(listMst.EmployeeID)

	list := DetailRFID{}
	err = db.DB.QueryRow("select RFIDID,EP.EmployeeID,EmployeeNumber,EmployeeName, "+
		"Plant as LocationCode,[Group] as GroupCode,Unit as UnitCode,1 as ProdCapacity,1 as ProdTarget "+
		"from [SKTIS].[dbo].[ExePlantWorkerAbsenteeism] EP "+
		"INNER JOIN [SKTIS].[dbo].[MstRFID] MR "+
		"ON EP.EmployeeID=MR.EmployeeID "+
		"Where EP.EmployeeID=$2 and IsActive=1 "+
		"and StartDateAbsent>=$1 and EndDateAbsent<=$1 "+
		"UNION ALL "+
		"SELECT RFIDID,PE.EmployeeID,EmployeeNumber,EmployeeName, Plant as LocationCode, "+
		"[Group] as GroupCode,Unit as UnitCode, COALESCE(ProdCapacity,0) AS ProdCapacity, COALESCE(ProdTarget,0) AS ProdTarget  "+
		"FROM [SKTIS].[dbo].[ExePlantProductionEntryVerification] PV "+
		"INNER JOIN [SKTIS].[dbo].[ExePlantProductionEntry] PE "+
		"ON PV.ProductionEntryCode=PE.ProductionEntryCode "+
		"INNER JOIN [SKTIS].[dbo].[MstRFID] MR "+
		"ON PE.EmployeeID=MR.EmployeeID "+
		"WHERE PE.EmployeeID=$2 and IsActive=1 "+
		"AND ProductionDate=$1 ", payload.Date, listMst.EmployeeID).Scan(&list.RFIDID, &list.EmployeeID,
		&list.EmployeeNumber, &list.EmployeeName, &list.LocationCode, &list.GroupCode,
		&list.UnitCode, &list.ProdCapacity, &list.ProdTarget)
	if err != nil {
		// data not found
		if payload.AbsentType == "ProductionTarget" {
			res.Message = "Data not found !"
			res.Status = "error"
			return
		}
	} else {
		if list.ProdTarget != 0 {
			res.Message = "Data exist !"
			res.Status = "error"
			return
		}
	}
	res.Message = "Data was successfully validated !"
	res.Status = "success"
	det := DetailRFID{}
	det.RFIDID = listMst.RFIDID
	det.EmployeeID = listMst.EmployeeID
	det.EmployeeNumber = listMst.EmployeeNumber
	det.EmployeeName = listMst.EmployeeName
	det.LocationCode = listMst.LocationCode
	det.GroupCode = listMst.GroupCode
	det.UnitCode = listMst.UnitCode
	det.ProdCapacity = list.ProdCapacity
	res.Data = det
	return
}
