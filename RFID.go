package main

import (
	// "fmt"

	ctl "sktisrfid/pkg/controllers"
	db "sktisrfid/pkg/database"

	"github.com/wailsapp/wails"
)

type RFID struct {
	message string
	runtime *wails.Runtime
	logger  *wails.CustomLogger
}

// NewTodos attempts to create a new Todo list
func NewRFID() (*RFID, error) {
	// Create new Todos instance
	result := &RFID{}
	// Return it
	return result, nil
}

func (t *RFID) GetListProductionTarget(data map[string]interface{}) []ctl.ListProductionTarget {
	return ctl.ListProdTarget(data)
}

func (t *RFID) GetListAbsenteeism(data map[string]interface{}) []ctl.ListAbsenteeism {
	return ctl.ListAbsent(data)
}

func (t *RFID) WailsInit(runtime *wails.Runtime) error {
	db.ConnectMySQL()
	t.runtime = runtime
	t.logger = t.runtime.Log.New("RFID")
	t.logger.Info("I'm here")

	t.runtime.Window.SetTitle("SKTISRFID")
	return nil
}

func (t *RFID) WailsShutdown() {
	defer db.DB.Close()
	// De-Allocate some resources...
}
