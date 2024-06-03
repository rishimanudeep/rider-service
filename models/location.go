package models

type RiderLocation struct {
	RiderID int `json:"rider_id"`
	Location
}

type Rider struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	VehicleDetails string `json:"vehicle_details"`
	Created        string `json:"created"`
	UpdatedAt      string `json:"updated_at"`
	Location
}

type Availability struct {
	IsAvailable bool `json:"is_available"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
