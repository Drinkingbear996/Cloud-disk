package models

import (
	"log"
	"xorm.io/xorm"
)
import _ "github.com/go-sql-driver/mysql"

var Engine = Init()

// Init initializes the database engine and returns an instance of *xorm.Engine.
func Init() *xorm.Engine {
	// Replace with your actual MySQL credentials and database information
	engine, err := xorm.NewEngine("mysql", "root:12345678@tcp(127.0.0.1:3306)/Cloud-data?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Xorm New Engine Error: %v", err)
		panic(err) // Optional: panic to prevent continuing if the DB connection fails
	}

	// Optionally, you can log the success of the engine initialization
	log.Println("Database engine successfully initialized")

	return engine
}
