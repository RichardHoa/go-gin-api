package main

import (
	// "fmt"
	"log"

	"github.com/RichardHoa/go-gin-api/cmd/api"
	"github.com/RichardHoa/go-gin-api/cmd/config"
	"github.com/RichardHoa/go-gin-api/cmd/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	// Create database
	db, createSQLErr := db.NewMySQLDB(mysql.Config{
		User:                 config.ENVs.DBUser,
		Passwd:               config.ENVs.DBPassword,
		Addr:                 config.ENVs.DBAddress,
		DBName:               config.ENVs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if createSQLErr != nil {
		log.Fatal(createSQLErr)
	}

	// Connect to database
	connectDBErr := db.Ping()
	if connectDBErr != nil {
		log.Fatal(connectDBErr)
	}
	log.Println("Connected to database")

	// Put DB connection to server
	port := ":" + config.ENVs.Port
	server := api.NewAPIServer(port, db)

	// Run server
	if serverRunErr := server.Run(); serverRunErr != nil {
		log.Fatal(serverRunErr)
	}

}

