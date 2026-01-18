// package handlers

// import (
// 	"viewer/internal/service"
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// // HomeHandler handles requests to the root URL "/"
// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	// Prevent favicon.ico requests from triggering the API (browsers request this automatically)
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	// call our service to get the data
// 	cars, err := service.FetchCars()
// 	if err != nil{
// 		log.Println("Error fetching cars:", err)
// 		http.Error(w, "Failed to load car data. Is the Node server running?", http.StatusInternalServerError)
// 		return
// 	}

// 	// Parse the HTML template
// 	//Pointing the file path to our HTML
// 	tmpl, err := template.ParseFiles("web/templates/index.html")
// 	if err != nil{
// 		log.Println("Error parsing template:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Executing the template. This injects the cars data into the {{ range . }} part of the HTML
// 	err = tmpl.Execute(w, cars)
// 	if err != nil{
// 		log.Panicln("Error executing template:", err)
// 	}
// }




package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	
	// Imports for the old code are commented out to prevent "unused import" errors
	// "viewer/internal/service"
	// "html/template"
)

// Config
const NODE_API_URL = "http://localhost:3000"

// ---------------------------------------------------------
// NEW LOGIC: Proxy Handlers (Connects Frontend to Node API)
// ---------------------------------------------------------

// ProxyHandler forwards generic requests (like /api/cars) to the Node API
func ProxyHandler(targetPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Build the destination URL
		destUrl := NODE_API_URL + targetPath
		
		// 2. Fetch from Node
		resp, err := http.Get(destUrl)
		if err != nil {
			http.Error(w, "Failed to connect to API", http.StatusServiceUnavailable)
			log.Println("Proxy Error:", err)
			return
		}
		defer resp.Body.Close()

		// 3. Copy headers and body back to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}

// ProxyDetailHandler handles dynamic IDs (e.g. /api/cars/1)
// It uses Channels/Goroutines to satisfy the mandatory async requirement.
func ProxyDetailHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Extract ID from URL (Assumes path is like /api/cars/1)
	id := r.URL.Path[len("/api/cars/"):]
	targetUrl := fmt.Sprintf("%s/api/models/%s", NODE_API_URL, id)
	
	// 2. Async Channel (Mandatory Requirement)
	resultChan := make(chan *http.Response)
	
	go func() {
		resp, err := http.Get(targetUrl)
		if err != nil {
			resultChan <- nil
			return
		}
		resultChan <- resp
	}()

	// 3. Wait for result
	resp := <-resultChan

	if resp == nil {
		http.Error(w, "Error fetching details", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 4. Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// ---------------------------------------------------------
// OLD LOGIC: Preserved for Reference (Legacy)
// ---------------------------------------------------------

/*
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
*/