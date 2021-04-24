package database

import (
	"database/sql"
	"github.com/bladedevleoper/go-cli-app/config"
	"github.com/bladedevleoper/go-cli-app/util"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

//construct the database credentials
func DbConnect() *sql.DB {
	util.LoadConfig()
	connection := config.DatabaseConfig{}
	connection.DB_DATABASE = os.Getenv("DB_DATABASE")
	connection.DB_HOST = os.Getenv("DB_HOST")
	connection.DB_CONNECTION = os.Getenv("DB_CONNECTION")
	connection.DB_PORT = os.Getenv("DB_PORT")
	connection.DB_USERNAME = os.Getenv("DB_USERNAME")
	connection.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	db, err := sql.Open(connection.DB_CONNECTION, getConnectionString(connection))

	if err != nil {
		panic(err.Error())
	}

	return db
}

func getConnectionString(connection config.DatabaseConfig) string {
	return connection.DB_USERNAME + ":" + "@tcp(" + connection.DB_HOST + ":" + connection.DB_PORT + ")/" + connection.DB_DATABASE
}
