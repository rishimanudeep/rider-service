package service

import (
	"errors"
	"github.com/rider/models"
)

type service struct {
	store riderService
}

// New unexported return type for exported function
// used factory for injecting dependencies
func New(s riderService) service {
	return service{store: s}
}

// UpdateRiderAvailability updates the availability of rider
func (s service) UpdateRiderAvailability(availability *models.Availability, riderID int) error {
	err := s.store.UpdateRiderAvailability(availability, riderID)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// UpdateRiderLocation updates the location in the DB by calling store layer
func (s service) UpdateRiderLocation(location *models.RiderLocation) error {
	err := s.store.UpdateRiderLocation(location)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s service) GetNearbyRiders(latitude, longitude float64, radius int) ([]models.RiderLocation, error) {
	resp, err := s.store.GetNearbyRiders(latitude, longitude, radius)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return resp, nil
}

// RegisterRiders call store layer to register and update Rider Location table
func (s service) RegisterRiders(rider models.Rider) (id int, err error) {
	id, err = s.store.RegisterRiders(rider)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	riderLocation := &models.RiderLocation{
		RiderID: id,
		Location: models.Location{
			Latitude:  rider.Latitude,
			Longitude: rider.Longitude,
		},
	}

	err = s.store.InsertRiderLocation(riderLocation)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	return id, nil
}

// UpdateRiderDetails call store layer to update details
func (s service) UpdateRiderDetails(rider models.Rider) error {
	return s.store.UpdateRiderDetails(rider)
}

// GetRiderDetails call store layer to getRider details
func (s service) GetRiderDetails(id int) (*models.Rider, error) {
	return s.store.GetRiderDetails(id)
}
