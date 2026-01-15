package models

// Specification represents the technical details nested inside a car object
type Specification struct {
	Engine			string			`json:"engine"`
	Horsepower		int				`json:"horsepower"`
	Transmission	string			`json:"transmission"`
	Drivetrain		string			`json:"drivetrain"`
}

// Car represents a single car object from the API
type Car struct {
	ID				int				`json:"id"`
	Name			string			`json:"name"`
	ManufacturerID	int				`json:"manufacturerId"`
	CategoryID		int				`json:"categoryId"`
	Year			int				`json:"year"`
	Specs			Specification	`json:"Specifications"`
	Price			string			`json:"price"`
	Availability	bool			`json:"availability"`
	Image			string			`json:"image"`
}

// Manufacturer matches the "manufacturers" array in our API
type Manufacturer struct {
	ID				int				`json:"id"`
	Name			string			`json:"name"`
	Country			string			`json:"country"`
	FoundingYear	int				`json:"foundingYear"`
}

// Category matches the "categories" array in our API
type Category struct {
	ID				int				`json:"id"`
	Name			string			`jsnon:"name"`
}

// ApiError represents an error response structure 
type ApiError struct {
	Error string `json:"error"`
}