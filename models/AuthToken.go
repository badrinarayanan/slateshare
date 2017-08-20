package models

import (

	"github.com/jinzhu/gorm"
)

/**
	Table for Auth Token
 */

type AuthToken struct {

	gorm.Model
	Token string

}
