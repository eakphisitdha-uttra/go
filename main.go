package main

import (
	"clean_architecture/databases/mongodb"
	"clean_architecture/databases/mssql"
	"clean_architecture/databases/mysql"
	"clean_architecture/databases/postgresql"
	"clean_architecture/logs"
	"clean_architecture/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// logs initial
	logs.InitLoggerError()
	defer logs.CloseLogError()

	logs.InitLoggerInfo()
	defer logs.CloseLogInfo()

	logs.InitLoggerDebug()
	defer logs.CloseLogDebug()

	// database connection
	pg, err := postgresql.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer postgresql.Close()

	mg, err := mongodb.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer mongodb.Close()

	my, err := mysql.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	ms, err := mssql.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer mssql.Close()

	_ = my
	_ = ms

	// set up routes
	r := routes.SetupRouter(pg, mg)
	// running
	r.Run(":8080")
}
