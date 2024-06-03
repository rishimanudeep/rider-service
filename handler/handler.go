package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rider/errors"
	"github.com/rider/models"
	"net/http"
	"strconv"
)

type handler struct {
	svc riderHandler
}

func New(h riderHandler) handler {
	return handler{svc: h}
}

func (h handler) UpdateRiderAvailability(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	riderID, err := strconv.Atoi(params["riderid"])
	if err != nil {
		http.Error(w, "Invalid rider ID", http.StatusBadRequest)
		return
	}

	var availability *models.Availability

	if err := json.NewDecoder(r.Body).Decode(&availability); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.svc.UpdateRiderAvailability(availability, riderID)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RiderLocation availability updated successfully"))
}

func (h handler) UpdateRiderLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	riderID, err := strconv.Atoi(params["riderid"])
	if err != nil {
		http.Error(w, "Invalid rider ID", http.StatusBadRequest)
		return
	}

	var location models.RiderLocation

	location.RiderID = riderID

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.svc.UpdateRiderLocation(&location)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RiderLocation location updated successfully"))
}

func (h handler) GetNearbyRiders(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	latitude, err := strconv.ParseFloat(queryValues.Get("latitude"), 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(queryValues.Get("longitude"), 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.Atoi(queryValues.Get("radius"))
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	resp, err := h.svc.GetNearbyRiders(latitude, longitude, radius)
	if err != nil {
		h.handleError(w, err)

		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h handler) RegisterRiders(w http.ResponseWriter, r *http.Request) {
	var rider models.Rider

	if err := json.NewDecoder(r.Body).Decode(&rider); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, err := h.svc.RegisterRiders(rider)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "RiderLocation registered successfully",
	})
}

func (h handler) UpdateRiderDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid rider ID", http.StatusBadRequest)
		return
	}

	var rider models.Rider

	if err = json.NewDecoder(r.Body).Decode(&rider); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	rider.ID = id

	err = h.svc.UpdateRiderDetails(rider)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "RiderLocation updated successfully",
	})
}

func (h handler) GetRiderDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid rider ID", http.StatusBadRequest)
		return
	}

	resp, err := h.svc.GetRiderDetails(id)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *errors.EntityNotFound:
		http.Error(w, e.Error(), http.StatusNotFound)
	case *errors.NoResponse:
		http.Error(w, e.Error(), http.StatusNotFound)
	case *errors.MissingParam:
		http.Error(w, e.Error(), http.StatusBadRequest)
	case *errors.ValidationError:
		http.Error(w, e.Error(), http.StatusBadRequest)
	case *errors.InternalServerError:
		http.Error(w, e.Error(), http.StatusInternalServerError)
	default:
		http.Error(w, "unknown error", http.StatusInternalServerError)
	}
}
