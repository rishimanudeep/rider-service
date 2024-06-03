package store

import (
	"database/sql"

	"github.com/rider/errors"
	"github.com/rider/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{
		db: db,
	}
}

func (s store) UpdateRiderAvailability(availability *models.Availability, riderID int) error {
	// Update rider availability in the database
	_, err := s.db.Exec("UPDATE rider_availability SET available = $1 WHERE rider_id = $2", availability.IsAvailable, riderID)
	if err != nil {
		return &errors.InternalServerError{Message: "Query Execution Error"}
	}

	return nil
}

func (s store) UpdateRiderLocation(location *models.RiderLocation) error {
	// Update rider location in the database
	_, err := s.db.Exec("UPDATE rider_availability SET latitude = $1, longitude = $2 WHERE rider_id = $3",
		location.Latitude, location.Longitude, location.RiderID)
	if err != nil {
		return &errors.InternalServerError{Message: "Query Execution Error"}
	}

	return nil
}

func (s store) GetNearbyRiders(latitude, longitude float64, radius int) ([]models.RiderLocation, error) {
	// Query nearby riders from the database
	rows, err := s.db.Query(`
			SELECT id, latitude, longitude
			FROM rider_availability
			WHERE ST_DWithin(
				ST_SetSRID(ST_MakePoint(longitude::double precision, latitude::double precision), 4326),
				ST_SetSRID(ST_MakePoint($1, $2), 4326),
				$3
			)
		`, longitude, latitude, radius)
	if err != nil {
		return nil, &errors.InternalServerError{Message: "Query Execution Error"}
	}
	defer rows.Close()

	var riders []models.RiderLocation
	for rows.Next() {
		var rider models.RiderLocation
		if err := rows.Scan(&rider.RiderID, &rider.Latitude, &rider.Longitude); err != nil {
			return nil, &errors.InternalServerError{Message: "DB Scan Err"}

		}
		riders = append(riders, rider)
	}

	if err := rows.Err(); err != nil {
		return nil, &errors.InternalServerError{Message: "rows error"}
	}

	return riders, nil
}

func (s store) InsertRiderLocation(rider *models.RiderLocation) error {
	query := `INSERT INTO rider_availability (rider_id,available,latitude,longitude) 
		          VALUES ($1,true,$2, $3) 
		          RETURNING id`
	var id int

	err := s.db.QueryRow(query, rider.RiderID, rider.Latitude, rider.Longitude).Scan(&id)
	if err != nil {
		return &errors.InternalServerError{Message: "Query Execution Error"}

	}

	return nil
}

func (s store) RegisterRiders(rider models.Rider) (int, error) {
	query := `INSERT INTO riders (name, email, vehicle_details, created, updated_at) 
		          VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
		          RETURNING id`
	var id int

	err := s.db.QueryRow(query, rider.Name, rider.Email, rider.VehicleDetails).Scan(&id)
	if err != nil {
		return 0, &errors.InternalServerError{Message: "Query Execution Error"}
	}

	return id, nil
}

func (s store) UpdateRiderDetails(rider models.Rider) error {
	query := `UPDATE riders SET 
				  name = COALESCE(NULLIF($1, ''), name), 
				  email = COALESCE(NULLIF($2, ''), email), 
				  vehicle_details = COALESCE(NULLIF($3, ''), vehicle_details), 
				  updated_at = CURRENT_TIMESTAMP 
				  WHERE id = $4`
	_, err := s.db.Exec(query, rider.Name, rider.Email, rider.VehicleDetails, rider.ID)
	if err != nil {
		return &errors.InternalServerError{Message: "Query Execution Error"}
	}

	return nil
}

func (s store) GetRiderDetails(id int) (*models.Rider, error) {
	var rider models.Rider

	query := `SELECT id, name, email, vehicle_details, created, updated_at FROM riders WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&rider.ID, &rider.Name, &rider.Email, &rider.VehicleDetails, &rider.Created, &rider.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &errors.EntityNotFound{"riders not found in db"}
		}

		return nil, &errors.InternalServerError{Message: "Query Execution Error"}
	}

	return &rider, nil
}
