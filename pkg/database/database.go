package database

import (
	// Configs
	"database/sql"
	"fmt"
	"log"
	cfg "sktisrfid/pkg/configs"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	DB *sql.DB
)

func ConnectMySQL() {
	dsn := cfg.GetConfig().GetMyConnectionInfo()
	db, err := sql.Open("mssql", dsn)

	if err != nil {
		fmt.Println(strings.Repeat("!", 40))
		fmt.Println("☹️  Could Not Establish MysQL DB Connection")
		fmt.Println(strings.Repeat("!", 40))
		log.Fatal(err)
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("😀 Connected To MysQL DB")
	fmt.Println(strings.Repeat("-", 40))

	DB = db
}
