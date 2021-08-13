package controllers

import (
	"log"
	db "sktisrfid/pkg/database"
	"time"
)

type PayloadProductionTarget struct {
	ProductionDate string
}

type ListProductionTarget struct {
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

func ListProdTarget(data map[string]interface{}) (res []ListProductionTarget) {
	var payload PayloadProductionTarget
	payload.ProductionDate = data["ProductionDate"].(string)

	rows, err := db.DB.Query("SELECT RFIDID,PE.EmployeeID,EmployeeNumber,EmployeeName, Plant as LocationCode, "+
		"[Group] as GroupCode,Unit as UnitCode, ProdCapacity, ProdTarget, PE.UpdatedDate as CreatedDate "+
		"FROM ExePlantProductionEntryVerification PV "+
		"INNER JOIN ExePlantProductionEntry PE "+
		"ON PV.ProductionEntryCode=PE.ProductionEntryCode "+
		"INNER JOIN MstRFID MR "+
		"ON PE.EmployeeID=MR.EmployeeID "+
		"WHERE IsFromRFID=1 and IsActive=1 "+
		"AND ProductionDate=$1 ", payload.ProductionDate)
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		list := ListProductionTarget{}
		err = rows.Scan(&list.RFIDID, &list.EmployeeID, &list.EmployeeNumber, &list.EmployeeName,
			&list.LocationCode, &list.GroupCode, &list.UnitCode, &list.ProdCapacity, &list.ProdTarget, &list.CreatedDate)
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
