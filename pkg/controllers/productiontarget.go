package controllers

import (
	"encoding/json"
	"fmt"
	db "sktisrfid/pkg/database"
	"time"
)

type PayloadUpdateDelete struct {
	update []PayloadProductionTarget
	delete []PayloadProductionTarget
}

type PayloadProductionTarget struct {
	ProductionDate string
	CreatedDate    string
	RFIDID         string
	EmployeeID     string
	ProdTarget     float32
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
		"[Group] as GroupCode,Unit as UnitCode, COALESCE(ProdCapacity,0) AS ProdCapacity, COALESCE(ProdTarget,0) AS ProdTarget, PE.UpdatedDate as CreatedDate "+
		"FROM [SKTIS].[dbo].[ExePlantProductionEntryVerification] PV "+
		"INNER JOIN [SKTIS].[dbo].[ExePlantProductionEntry] PE "+
		"ON PV.ProductionEntryCode=PE.ProductionEntryCode "+
		"INNER JOIN [SKTIS].[dbo].[MstRFID] MR "+
		"ON PE.EmployeeID=MR.EmployeeID "+
		"WHERE IsFromRFID=1 and IsActive=1 "+
		"AND ProductionDate=$1 ORDER BY PE.UpdatedDate DESC", payload.ProductionDate)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}
	for rows.Next() {
		list := ListProductionTarget{}
		err = rows.Scan(&list.RFIDID, &list.EmployeeID, &list.EmployeeNumber, &list.EmployeeName,
			&list.LocationCode, &list.GroupCode, &list.UnitCode, &list.ProdCapacity, &list.ProdTarget, &list.CreatedDate)
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

func UpdateDeleteRFIDProductionTarget(data map[string]interface{}) (res ResponseResult) {
	fmt.Println(data)
	var payload PayloadUpdateDelete
	ParamUpdate, _ := json.Marshal(data["update"])
	if err := json.Unmarshal(ParamUpdate, &payload.update); err != nil {
		// log.Fatal(err)
		res.Status = "error"
		res.Message = err.Error()
		return
	}

	ParamDelete, _ := json.Marshal(data["delete"])
	if err := json.Unmarshal(ParamDelete, &payload.delete); err != nil {
		// log.Fatal(err)
		res.Status = "error"
		res.Message = err.Error()
		return
	}

	if payload.delete != nil {
		for i := 0; i < len(payload.delete); i++ {

			_, err := db.DB.Query("EXEC [SKTIS].[dbo].[UPDATE_PRODUCTION_ENTRY_RFID] @EmployeeID = $1 ,@ProdTarget = $2, @CreatedDate = $3, @ProductionDate = $4, @UpdatedBy = $5, @Mode = $6", payload.delete[i].EmployeeID, payload.delete[i].ProdTarget, payload.delete[i].CreatedDate, payload.delete[i].ProductionDate, "RFID", "DELETE")

			if err != nil {
				// log.Fatal(err)
				res.Status = "error"
				res.Message = err.Error()
				return
			}

		}
	}

	if payload.update != nil {
		for i := 0; i < len(payload.update); i++ {
			_, err := db.DB.Query("EXEC [SKTIS].[dbo].[UPDATE_PRODUCTION_ENTRY_RFID] @EmployeeID = $1 ,@ProdTarget = $2, @CreatedDate = $3, @ProductionDate = $4, @UpdatedBy = $5, @Mode = $6", payload.update[i].EmployeeID, payload.update[i].ProdTarget, payload.update[i].CreatedDate, payload.update[i].ProductionDate, "RFID", "UPDATE")

			if err != nil {
				// log.Fatal(err)
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
