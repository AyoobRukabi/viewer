package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"viewer/internal/models"
)

// API url

const apiUrl = "http://localhost:3000"

// FetchCars gets the list of all cars models
func FetchCars() ([]models.Car, error) {

	// the specific url
	url := fmt.Sprintf("%s/models", apiUrl)

	// Creating a Client with a timeout so our app doesn't freeze
	client := http.Client{
		Timeout: time.Second * 10,
	}

	// Making the Request
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	// Schedule the connection to close when this function finishes
	defer resp.Body.Close()

	// Check if the server said "OK" {status Code 200}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Decode directly into a slice (list) of Cars
	var cars []models.Car
	if err := json.NewDecoder(resp.Body).Decode(&cars); err != nil {
		return nil, err
	}

	// Return the list of cars
	return cars, nil

}