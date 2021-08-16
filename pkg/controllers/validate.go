package controllers

import (
	"database/sql"
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
		"FROM MstRFID WHERE RFIDID=$1", payload.RFIDID).Scan(&listMst.RFIDID, &listMst.EmployeeID, &listMst.EmployeeNumber,
		&listMst.EmployeeName, &listMst.LocationCode, &listMst.UnitCode, &listMst.GroupCode)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		res.Message = "Data not found !"
		res.Status = "error"
		res.Data = DetailRFID{}
		return
	}

	list := DetailRFID{}
	err = db.DB.QueryRow("select EP.EmployeeID,EmployeeNumber,EmployeeName, "+
		"Plant as LocationCode,[Group] as GroupCode,Unit as UnitCode,null as ProdCapacity,null as ProdTarget "+
		"from ExePlantWorkerAbsenteeism EP "+
		"Where IsFromRFID=1 and EP.EmployeeID=$2 "+
		"and AbsentType='Alpa' and StartDateAbsent>=$1 and EndDateAbsent<=$1 "+
		"UNION ALL "+
		"SELECT PE.EmployeeID,EmployeeNumber,EmployeeName, Plant as LocationCode, "+
		"[Group] as GroupCode,Unit as UnitCode, ProdCapacity, ProdTarget "+
		"FROM ExePlantProductionEntryVerification PV "+
		"INNER JOIN ExePlantProductionEntry PE "+
		"ON PV.ProductionEntryCode=PE.ProductionEntryCode "+
		"WHERE IsFromRFID=1 and EP.EmployeeID=$2 "+
		"AND ProductionDate=$1 ", payload.Date, listMst.EmployeeID).Scan(&list.RFIDID, &list.EmployeeID,
		&list.EmployeeNumber, &list.EmployeeName, &list.LocationCode, &list.GroupCode,
		&list.UnitCode, &list.ProdCapacity, &list.ProdTarget)
	if err != nil {
		// data not found
	} else {
		res.Message = "Data exist !"
		res.Status = "error"
		return
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
