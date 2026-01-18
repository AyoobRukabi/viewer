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
	"log"
	"net/http"
	"viewer/internal/handlers" 
)

const PORT = ":8080"

func main() {
	// 1. Serve Images
	imgFs := http.FileServer(http.Dir("./api/img"))
	http.Handle("/img/", http.StripPrefix("/img/", imgFs))

	// 2. Serve Website
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// 3. API Proxy Endpoints (Using the 'handlers' package)
	http.HandleFunc("/api/cars", handlers.ProxyHandler("/api/models"))
	http.HandleFunc("/api/cars/", handlers.ProxyDetailHandler)
	http.HandleFunc("/api/manufacturers", handlers.ProxyHandler("/api/manufacturers"))
	http.HandleFunc("/api/categories", handlers.ProxyHandler("/api/categories"))

	// 4. Start Server
	fmt.Printf("✅ Server running!\n")
	fmt.Printf("🌎 Website: http://localhost%s\n", PORT)
	
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}