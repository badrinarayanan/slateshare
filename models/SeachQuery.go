package models

import "github.com/jinzhu/gorm"

type SearchQuery struct {

	gorm.Model
	SearchParam string
}
