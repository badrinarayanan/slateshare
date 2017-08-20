package main

import(
	"net/http"
	"slateshare/controllers"
	"github.com/joho/godotenv"
	"slateshare/models"
	"log"
	"os"

)

func main(){


	mux := http.NewServeMux()

	// Routes

	mux.HandleFunc("/search",controllers.SearchTwitter)

	// Read the Env File

	err := godotenv.Load()
	if(err != nil){
		log.Fatal("Cannot load Env file")
	}
	if(os.Getenv("migrate") == "1"){

		models.MigrateDB()
	}

	// Start Server

	server := &http.Server{

		Addr:"0.0.0.0:8080",
		Handler:mux,
	}

	server.ListenAndServe()

}
