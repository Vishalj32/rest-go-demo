package mysql

import (
	"log"
	"rest-go-demo/entity"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

//Connect creates Mysql Database connection
func Connect(connectionString string) (*gorm.DB, error) {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	log.Println("Connection was successful!!")
	return Connector, nil
}

//MigrateDB migrating existing struct to DB
func MigrateDB(connector *gorm.DB, table *entity.Person) {
	connector.AutoMigrate(&table)
	log.Println("Table migrated!!")
}
