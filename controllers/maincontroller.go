package controllers

import (
	"net/http"
	"slateshare/twitter"

	"net/url"

	"encoding/json"

	"slateshare/repository"
)

func SearchTwitter(w http.ResponseWriter,r *http.Request){

	var params map[string]string
	defer r.Body.Close()
	decodeparams := json.NewDecoder(r.Body)
	err := decodeparams.Decode(&params)
	if(err != nil){

		panic("Error Occurred")
	}

	var token string
	get_token_database := twitter.GetToken()
	if(get_token_database != ""){

		token = get_token_database
	}else{
		token = twitter.Twitter()
	}
	client := &http.Client{}

	req,_ := http.NewRequest("GET","https://api.twitter.com/1.1/search/tweets.json?q="+ url.QueryEscape(params["search"])+"&result_type=recent",nil)
	req.Header.Add("Authorization","Bearer " +token)
	resp,err := client.Do(req)
	if(err != nil){

		panic("Error")
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response twitter.Response
	err = decoder.Decode(&response)
	if(err != nil){

		panic("Error occurred")
	}
	paramid := repository.InsertSearchParam(params["search"])

	data := repository.GetTwitterBlockQuote(response,paramid)

	jData,err := json.Marshal(data)

	if(err != nil){
		panic("Error occurred")
	}
	w.Write(jData)
	w.Header().Set("Content-Type","application/json")
}
