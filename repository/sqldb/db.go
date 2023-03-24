package sqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"love-date/config"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {
	conf := config.New()
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", conf.SqlDB.Username, conf.SqlDB.Password,
		conf.SqlDB.Host, conf.SqlDB.Port, conf.SqlDB.Name)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{db: db}
}
