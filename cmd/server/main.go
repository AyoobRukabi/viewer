// package main

// import (
// 	"log"
// 	"net/http"
// 	"viewer/internal/handlers"
// )
// //test
// func main() {
// 	// Serve Static Files (CSS, JS)
// 	// This tells GO: if a URL starts with /static/, look inside the web/static folder"
// 	fs := http.FileServer(http.Dir("./web/static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	// Register Routes. When someone visits "/", run the HomeHandler function
// 	http.HandleFunc("/", handlers.HomeHandler)

// 	// Starting the server
// 	log.Println("Starting server on http://localhost:8080")
// 	// This starts the server and blocks forever. If it crashes it returns an error.
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Config
const (
	PORT         = ":8080"
	NODE_API_URL = "http://localhost:3000"
)

func main() {
	// 1. SERVE IMAGES (Clean URL: /img/car.jpg)
	// We map the URL path "/img/" to the folder "./api/img"
	imgFs := http.FileServer(http.Dir("./api/img"))
	http.Handle("/img/", http.StripPrefix("/img/", imgFs))

	// 2. SERVE WEBSITE (Clean URL: / )
	// We map the Root URL "/" to the folder "./frontend"
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// 3. API PROXY ENDPOINTS
	http.HandleFunc("/api/cars", proxyHandler("/api/models"))
	http.HandleFunc("/api/cars/", proxyDetailHandler)
	http.HandleFunc("/api/manufacturers", proxyHandler("/api/manufacturers"))
	http.HandleFunc("/api/categories", proxyHandler("/api/categories"))

	// 4. START SERVER
	fmt.Printf("✅ Server running!\n")
	fmt.Printf("🌎 Website: http://localhost%s\n", PORT) // Clickable link to root
	
	// Better error handling for port conflicts
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("❌ Server failed to start. Is port 8080 already in use?", err)
	}
}

// --- KEEP THE EXISTING PROXY FUNCTIONS BELOW (No changes needed there) ---

func proxyHandler(targetPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		destUrl := NODE_API_URL + targetPath
		resp, err := http.Get(destUrl)
		if err != nil {
			http.Error(w, "Failed to connect to API", http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}

func proxyDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/cars/"):]
	targetUrl := fmt.Sprintf("%s/api/models/%s", NODE_API_URL, id)
	
	resultChan := make(chan *http.Response)
	go func() {
		resp, err := http.Get(targetUrl)
		if err != nil {
			resultChan <- nil
			return
		}
		resultChan <- resp
	}()

	resp := <-resultChan
	if resp == nil {
		http.Error(w, "Error fetching car details", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}