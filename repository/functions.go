package repository

import (

	"slateshare/database"
	"slateshare/models"
	"slateshare/twitter"

)

/**
	Insert Search Parameter into Database
 */
func InsertSearchParam(params string) uint {

	db := database.Conn()
	var searchparam models.SearchQuery
	searchparam.SearchParam = params
	db.Create(&searchparam)
	return searchparam.ID
}

/**
	Insert Record into DB
 */
func InsertRecordDB(response models.TwitterData){

	db := database.Conn()
	db.Create(&response)
	database.CloseConnection(db)


}


/**
	Get Twitter BlockQoute
 */
func GetTwitterBlockQuote(response twitter.Response,searchid uint) []models.TwitterData {


	var blockquote []models.TwitterData

	for i:=0; i <  len(response.Statuses); i++ {

		url := "https://twitter.com/" + response.Statuses[i].User.ScreenName + "/status/" + response.Statuses[i].IDStr


		/** Create the model **/
		var details models.TwitterData
		details.FollowersCount = response.Statuses[i].User.FollowersCount
		details.TweetDate = response.Statuses[i].CreatedAt
		details.Location =  response.Statuses[i].User.Location
		details.UserName =  response.Statuses[i].User.Name
		details.ScreenName = response.Statuses[i].User.ScreenName
		details.TwitterId = response.Statuses[i].IDStr
		details.SearchId = searchid
		details.TwitterText = response.Statuses[i].Text
		details.ProfilePic = response.Statuses[i].User.ProfileImageURLHTTPS
		details.Description = response.Statuses[i].User.Description
		details.UserId = response.Statuses[i].User.IDStr
		details.TweetUrl = url
		/**  **/
		go InsertRecordDB(details)

		blockquote = append(blockquote,details)

	}

	return blockquote

}