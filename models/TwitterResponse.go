package models

import (

	"github.com/jinzhu/gorm"

)

// Table For Twitter Data

type TwitterData struct {


	gorm.Model
	SearchId uint
	TwitterId string
	TwitterText string
	ProfilePic string
	UserName string
	UserId string
	ScreenName string
	Location string
	Description string
	FollowersCount int
	TweetDate string
	TweetUrl string


}

