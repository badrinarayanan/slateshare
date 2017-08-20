package twitter

import (
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"os"
	"slateshare/database"
	"slateshare/models"
)

func Twitter() string{

	consumer_key := os.Getenv("consumer_key")
	secret_key   := os.Getenv("secret_key")

	consumer_key_escaped  := url.QueryEscape(consumer_key)
	secret_key_escaped    := url.QueryEscape(secret_key)

	final_key := consumer_key_escaped + ":" + secret_key_escaped
	bencode := base64.StdEncoding.EncodeToString([]byte(final_key))


	payload := strings.NewReader("grant_type=client_credentials")
	client := &http.Client{}

	req,_ := http.NewRequest("POST","https://api.twitter.com/oauth2/token",payload)

	req.Header.Add("Authorization","Basic "+bencode)
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if(err != nil){

		fmt.Println(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var t map[string]string
	err = decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	token := models.AuthToken{Token:t["access_token"]}
	db := database.Conn()
	db.Create(&token)
	database.CloseConnection(db)
	return t["access_token"]

}


func GetToken() string{

db := database.Conn()
var records models.AuthToken
db.First(&records)
database.CloseConnection(db)
return records.Token

}