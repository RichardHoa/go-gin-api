package main

import (
	// "database/sql"
	"log"
	"os"

	"github.com/RichardHoa/go-gin-api/cmd/config"
	"github.com/RichardHoa/go-gin-api/cmd/db"
	MySQLConfig "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	db, err := db.NewMySQLDB(MySQLConfig.Config{
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

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrate, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := migrate.Up(); err != nil {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := migrate.Down(); err != nil {
			log.Fatal(err)
		}
	}
}
