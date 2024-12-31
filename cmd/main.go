package main

import (
	// "fmt"
	"database/sql"
	"log"

	"github.com/RichardHoa/go-gin-api/cmd/api"
	"github.com/RichardHoa/go-gin-api/cmd/config"
	"github.com/RichardHoa/go-gin-api/cmd/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLDB(mysql.Config{
		User:                 config.ENVs.DBUser,
		Passwd:               config.ENVs.DBPassword,
		Addr:                 config.ENVs.DBAddress,
		DBName:               config.ENVs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	dbErr := initDB(db)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	log.Println("Connected to database")

	return nil

}
