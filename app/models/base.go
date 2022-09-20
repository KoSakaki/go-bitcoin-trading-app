package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gotrading/config"

	_ "github.com/mattn/go-sqlite3"
)

const tableNameSignalEvents = "Signal_events"

var DbConnection *sql.DB

func GetCandoleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY NOT NULL,
			product_code STRING,
			side STRING,
			price FLOAT,
			size FLOAT
		)`, tableNameSignalEvents)

	if _, err := DbConnection.Exec(cmd); err != nil {
		log.Println(err)
	}
	for _, duration := range config.Config.Durations {
		tableName := GetCandoleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				time DATETIME PRIMARY KEY NOT NULL,
				open FLOAT,
				close FLOAT,
				high FLOAT,
				low FLOAT,
				volume FLOAT
			)`, tableName)
		if _, err := DbConnection.Exec(c); err != nil {
			log.Println(err)
		}
	}
}
