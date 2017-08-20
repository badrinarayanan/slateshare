package database

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	
)

var db *gorm.DB

/**
Connection function
 */

func Conn() *gorm.DB{

	db, _= gorm.Open("postgres", "host=localhost user=badrinarayanananravi dbname=slateshare sslmode=disable password=")

	return db

}


func CloseConnection(db *gorm.DB){

	db.Close()
}
