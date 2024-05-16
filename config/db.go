package config

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() *sqlx.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := sqlx.Open("mysql", uri)
	if err != nil {
		log.Fatalf("Can't connect to database with err: %s", err)
	}

	keepAliveInterval, _ := time.ParseDuration("1m")

	go keepAlive(db, "mysql", dbname, keepAliveInterval)

	fmt.Println("Successfully connect to database")

	return db
}

func keepAlive(db *sqlx.DB, driver, schema string, interval time.Duration) {
	for {
		err := db.Ping()
		if err != nil {
			log.Printf("ERROR db.keepAlive driver=%s schema=%s \n%s \n\n", driver, schema, err)
		}

		time.Sleep(interval)
	}
}
