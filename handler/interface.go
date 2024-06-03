package handler

import (
	"github.com/rider/models"
)

type riderHandler interface {
	UpdateRiderAvailability(availability *models.Availability, riderID int) error
	UpdateRiderLocation(location *models.RiderLocation) error
	GetNearbyRiders(latitude, longitude float64, radius int) ([]models.RiderLocation, error)
	RegisterRiders(rider models.Rider) (id int, err error)
	UpdateRiderDetails(rider models.Rider) error
	GetRiderDetails(id int) (*models.Rider, error)
}
