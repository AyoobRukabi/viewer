package handlers

import (
	"viewer/internal/service"
	"html/template"
	"log"
	"net/http"
)

// HomeHandler handles requests to the root URL "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Prevent favicon.ico requests from triggering the API (browsers request this automatically)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// call our service to get the data
	cars, err := service.FetchCars()
	if err != nil{
		log.Println("Error fetching cars:", err)
		http.Error(w, "Failed to load car data. Is the Node server running?", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template
	//Pointing the file path to our HTML
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil{
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Executing the template. This injects the cars data into the {{ range . }} part of the HTML
	err = tmpl.Execute(w, cars)
	if err != nil{
		log.Panicln("Error executing template:", err)
	}
}