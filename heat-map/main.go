package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type model struct {
	lat float64
	lon float64
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := model{
		lat: 11.0148,
		lon: 76.965,
	}
	tpl.Execute(w, p)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/get-data", func(w http.ResponseWriter, r *http.Request) {
		// Your Go function that generates data
		// data := getData()

		// Encode the data as JSON and write it to the response
		// json.NewEncoder(w).Encode(data)
	})
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatalln("Error: File could not be opened")

	}

}
