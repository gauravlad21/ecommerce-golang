package dbhelper

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

var dbOps DbOperationsIF
var db *sql.DB

func GetDb() *sql.DB {
	if db == nil {
		db = connectDatabaseSql()
	}
	return db
}

func GetDbOps() DbOperationsIF {
	if dbOps == nil {
		dbOps = New(GetDb())
	}
	return dbOps
}

func connectDatabaseSql() *sql.DB {
	dbUrl := GetDbUrl()
	db := "postgres"
	sqlxDb, err := sql.Open(db, dbUrl)
	if err != nil {
		fmt.Printf("dbUrl: %v\n", dbUrl)
		panic(err)
	}
	err = sqlxDb.Ping()
	if err != nil {
		panic(err)
	}

	setDbPool(sqlxDb)
	return sqlxDb
}

func setDbPool(sqlDb *sql.DB) {
	maxIdleCon := viper.GetInt("database.max_idle_conn")
	if maxIdleCon != 0 {
		sqlDb.SetMaxIdleConns(maxIdleCon) // defaultMaxIdleConns = 2
	}

	sqlDb.SetMaxOpenConns(viper.GetInt("database.max_open_conn")) // The default is 0 (unlimited)
	sqlDb.SetConnMaxLifetime(0)                                   // 0, connections are reused forever.

}

func GetDbUrl() string {
	// "host=localhost user=gauravlad password=abc@123@ABC dbname=goskel port=5432 sslmode=disable"
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	username := viper.GetString("database.username")
	databaseName := viper.GetString("database.dbname")
	password := viper.GetString("database.password")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
}
