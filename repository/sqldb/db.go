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

var conf *config.Config

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
	conf = config.New()

	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", conf.SqlDB.Username, conf.SqlDB.Password,
		conf.SqlDB.Host, conf.SqlDB.Port, conf.SqlDB.Name)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}
	initDB(db)

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{db: db}
}

func initDB(db *sql.DB) {
	//ctx := context.Background()

	//_, cErr := db.Exec("create database if not exists " + conf.SqlDB.Name)
	//if cErr != nil {
	//	panic(fmt.Errorf("Error %w when creating DB\n", cErr))
	//}
	//_, err := db.ExecContext(ctx, "USE "+conf.SqlDB.Name)
	//if err != nil {
	//	panic(err)
	//}

	createUserTableQuery := `CREATE TABLE IF NOT EXISTS users(
  id INT NOT NULL AUTO_INCREMENT,
  email VARCHAR(45) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
  UNIQUE INDEX email_UNIQUE (email ASC) VISIBLE)`

	_, cErr := db.Exec(createUserTableQuery)
	if cErr != nil {
		panic(fmt.Errorf("Error %w when creating DB\n", cErr))
	}

	createProfileTableQuery := `CREATE TABLE IF NOT EXISTS profiles(
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  name VARCHAR(45) NULL,
  special_days_notify_active TINYINT NOT NULL,
  birthday_notify_active TINYINT NOT NULL,
  vip_active TINYINT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
  UNIQUE INDEX user_id_UNIQUE (user_id ASC) VISIBLE,
      FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE)`

	_, cErr = db.Exec(createProfileTableQuery)
	if cErr != nil {
		panic(fmt.Errorf("Error %w when creating DB\n", cErr))
	}

	createPartnerTableQuery := `CREATE TABLE IF NOT EXISTS  partners(
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  name VARCHAR(45) NULL,
  birthday DATETIME NULL,
  first_date DATETIME NOT NULL,
  is_deleted TINYINT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
   FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE)`
	_, cErr = db.Exec(createPartnerTableQuery)
	if cErr != nil {
		panic(fmt.Errorf("Error %w when creating DB\n", cErr))
	}

}
