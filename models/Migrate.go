package models

import (
	"slateshare/database"
)


func MigrateDB(){

	 db := database.Conn()
	 db.AutoMigrate(&AuthToken{},&SearchQuery{},&TwitterData{})
	 database.CloseConnection(db)

}



