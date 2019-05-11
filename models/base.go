package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"primary_key;type:varchar(36)"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid.String())
}

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, dbHost, dbPort, dbName) //Build connection string
	fmt.Println("======")
	fmt.Println("DB connection string", dbUri)
	fmt.Println("======")

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &DSRItem{}, &CustomerBill{}, &BillItem{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
