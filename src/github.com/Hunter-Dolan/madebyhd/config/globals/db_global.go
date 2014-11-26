package global

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseConnection struct {
	Connection gorm.DB
}

var DB *DatabaseConnection = nil

func InitDB() {
	DB = new(DatabaseConnection)
	DB.DBConnect()
}

// Private Methods

func (db_struct *DatabaseConnection) DBConnect() {
	//db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True")
	db, _ := gorm.Open("sqlite3", "./tmp/db.db")

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db_struct.Connection = db
}
// Public Methods

func (db_struct *DatabaseConnection) Debug() {
	db_struct.Connection.LogMode(true)
}
