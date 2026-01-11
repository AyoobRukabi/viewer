package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Car represents a car model with all its details
type Car struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	Manufacturer string                 `json:"manufacturer"`
	Category     string                 `json:"category"`
	Year         int                    `json:"year"`
	Price        int                    `json:"price"`
	ImageURL     string                 `json:"imageUrl"`
	Details      map[string]interface{} `json:"details,omitempty"`
}

// Manufacturer represents a car manufacturer
type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	FoundingYear int    `json:"foundingYear"`
	Logo         string `json:"logo"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// CarRequest represents a request for car data via channel
type CarRequest struct {
	ID       int
	Response chan CarResponse
}

// CarResponse represents the response sent through the channel
type CarResponse struct {
	Car   *Car
	Error error
}

var (
	// Mock data store
	cars          []Car
	manufacturers []Manufacturer
	// Channel for handling car detail requests asynchronously
	carRequestChan = make(chan CarRequest)
)

func init() {
	// Initialize manufacturers
	manufacturers = []Manufacturer{
		{
			ID:           1,
			Name:         "Audi",
			Country:      "Germany",
			FoundingYear: 1909,
			Logo:         "🅰️",
		},
		{
			ID:           2,
			Name:         "Mercedes-Benz",
			Country:      "Germany",
			FoundingYear: 1926,
			Logo:         "⭐",
		},
		{
			ID:           3,
			Name:         "BMW",
			Country:      "Germany",
			FoundingYear: 1916,
			Logo:         "🔵",
		},
		{
			ID:           4,
			Name:         "Toyota",
			Country:      "Japan",
			FoundingYear: 1937,
			Logo:         "🔴",
		},
	}

	// Initialize cars with specific test data
	cars = []Car{
		{
			ID:           1,
			Name:         "Audi A4",
			Manufacturer: "Audi",
			Category:     "Sedan",
			Year:         2024,
			Price:        42000,
			ImageURL:     "https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=400",
			Details: map[string]interface{}{
				"engine":       "2.0L Inline-4",
				"horsepower":   201,
				"transmission": "7-speed Automatic",
				"drivetrain":   "All-Wheel Drive",
			},
		},
		{
			ID:           2,
			Name:         "Mercedes-Benz E-Class",
			Manufacturer: "Mercedes-Benz",
			Category:     "Luxury Sedan",
			Year:         2024,
			Price:        62000,
			ImageURL:     "https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=400",
			Details: map[string]interface{}{
				"engine":       "2.0L Inline-4 Turbo",
				"horsepower":   255,
				"transmission": "9-speed Automatic",
				"drivetrain":   "Rear-Wheel Drive",
			},
		},
		{
			ID:           3,
			Name:         "BMW 3 Series",
			Manufacturer: "BMW",
			Category:     "Sport Sedan",
			Year:         2024,
			Price:        45000,
			ImageURL:     "https://images.unsplash.com/photo-1555215695-3004980ad54e?w=400",
			Details: map[string]interface{}{
				"engine":       "2.0L Inline-4 Turbo",
				"horsepower":   255,
				"transmission": "8-speed Automatic",
				"drivetrain":   "Rear-Wheel Drive",
			},
		},
		{
			ID:           4,
			Name:         "Audi Q5",
			Manufacturer: "Audi",
			Category:     "SUV",
			Year:         2024,
			Price:        48000,
			ImageURL:     "https://images.unsplash.com/photo-1609521263047-f8f205293f24?w=400",
			Details: map[string]interface{}{
				"engine":       "2.0L Inline-4 Turbo",
				"horsepower":   261,
				"transmission": "7-speed Automatic",
				"drivetrain":   "All-Wheel Drive",
			},
		},
		{
			ID:           5,
			Name:         "Toyota Camry",
			Manufacturer: "Toyota",
			Category:     "Sedan",
			Year:         2024,
			Price:        28000,
			ImageURL:     "https://images.unsplash.com/photo-1621007947382-bb3c3994e3fb?w=400",
			Details: map[string]interface{}{
				"engine":       "2.5L Inline-4",
				"horsepower":   203,
				"transmission": "8-speed Automatic",
				"drivetrain":   "Front-Wheel Drive",
			},
		},
		{
			ID:           6,
			Name:         "BMW X5",
			Manufacturer: "BMW",
			Category:     "SUV",
			Year:         2024,
			Price:        65000,
			ImageURL:     "https://images.unsplash.com/photo-1627454820516-1cf4b8f1ec30?w=400",
			Details: map[string]interface{}{
				"engine":       "3.0L Inline-6 Turbo",
				"horsepower":   335,
				"transmission": "8-speed Automatic",
				"drivetrain":   "All-Wheel Drive",
			},
		},
	}

	// Start the asynchronous car detail processor
	go carDetailProcessor()
}

// carDetailProcessor is a goroutine that processes car detail requests asynchronously
// This demonstrates the use of goroutines and channels for concurrent operations
func carDetailProcessor() {
	log.Println("Car detail processor started and waiting for requests...")

	for request := range carRequestChan {
		// Simulate some processing time (e.g., fetching from database or external API)
		time.Sleep(100 * time.Millisecond)

		// Find the car by ID
		var foundCar *Car
		for i := range cars {
			if cars[i].ID == request.ID {
				foundCar = &cars[i]
				break
			}
		}

		// Send response back through the channel
		if foundCar != nil {
			log.Printf("Processing completed for car ID %d: %s", request.ID, foundCar.Name)
			request.Response <- CarResponse{Car: foundCar, Error: nil}
		} else {
			log.Printf("Car ID %d not found", request.ID)
			request.Response <- CarResponse{Car: nil, Error: fmt.Errorf("car not found")}
		}
	}
}

// enableCORS adds CORS headers to allow frontend requests
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// sendJSON sends a JSON response
func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// sendError sends an error response
func sendError(w http.ResponseWriter, message string, code int) {
	response := ErrorResponse{
		Error:   http.StatusText(code),
		Message: message,
		Code:    code,
	}
	sendJSON(w, response, code)
}

// handleGetCars returns all cars
func handleGetCars(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("GET /api/cars - Fetching all cars")
	sendJSON(w, cars, http.StatusOK)
}

// handleGetCarByID returns a specific car using goroutines and channels
func handleGetCarByID(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/cars/")
	id, err := strconv.Atoi(path)
	if err != nil {
		sendError(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	log.Printf("GET /api/cars/%d - Requesting car details asynchronously", id)

	// Create a response channel for this specific request
	responseChan := make(chan CarResponse)

	// Send request to the car detail processor goroutine
	carRequestChan <- CarRequest{
		ID:       id,
		Response: responseChan,
	}

	// Wait for response from the goroutine
	response := <-responseChan
	close(responseChan)

	if response.Error != nil {
		sendError(w, "Car not found", http.StatusNotFound)
		return
	}

	sendJSON(w, response.Car, http.StatusOK)
}

// handleGetManufacturers returns all manufacturers
func handleGetManufacturers(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("GET /api/manufacturers - Fetching all manufacturers")
	sendJSON(w, manufacturers, http.StatusOK)
}

// handle404 handles not found errors
func handle404(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	log.Printf("404 Not Found: %s %s", r.Method, r.URL.Path)
	sendError(w, "The requested resource was not found", http.StatusNotFound)
}

// routeHandler is a custom router that handles all routes
func routeHandler(w http.ResponseWriter, r *http.Request) {
	// Recover from panics to send 500 errors
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Internal Server Error: %v", err)
			enableCORS(w)
			sendError(w, "Internal server error occurred", http.StatusInternalServerError)
		}
	}()

	path := r.URL.Path

	switch {
	case path == "/api/cars":
		handleGetCars(w, r)
	case strings.HasPrefix(path, "/api/cars/"):
		handleGetCarByID(w, r)
	case path == "/api/manufacturers":
		handleGetManufacturers(w, r)
	default:
		handle404(w, r)
	}
}

func main() {
	// Set up routes
	http.HandleFunc("/", routeHandler)

	// Serve static files (frontend)
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	port := "8080"
	log.Printf("Server starting on http://localhost:%s", port)
	log.Println("API Endpoints:")
	log.Println("  GET  /api/cars")
	log.Println("  GET  /api/cars/{id}")
	log.Println("  GET  /api/manufacturers")
	log.Println("")
	log.Println("Frontend available at: http://localhost:8080/frontend/")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
