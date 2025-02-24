package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-net_http-project/internal/config"
	"go-net_http-project/internal/dto/requests"
	"go-net_http-project/internal/helpers"
	"go-net_http-project/models"
	"log"
	"net/http"
	"strconv"
)

type EventHandler struct {
	*CommonHandlers
	validator *validator.Validate
}

func NewEventHandler(commonHandlers *CommonHandlers) *EventHandler {

	return &EventHandler{
		CommonHandlers: commonHandlers,
	}
}

// GetEvents 	@Summary List events
// @Description	Get all events
// @Summary		Events
// @Description	Returns user events
// @Tags		Events
// @Produce		json
// @Success		200	{string}	string	"OK"
// @Router		/events [get]
func (h *EventHandler) GetEvents(w http.ResponseWriter, _ *http.Request) {
	functionName := fmt.Sprintf("%s.%s",
		config.AppConfig.Schemas["Public"],
		config.AppConfig.DatabaseFunctions["GetEvents"])

	events, err := h.DbService.RunFunction(context.Background(), functionName, nil)
	if err != nil {
		http.Error(w, "Failed to get events", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"events":  events,
		"message": "Events retrieved successfully",
		"status":  http.StatusOK,
	})
}

// FindEvent 	@Summary Find event
// @Description	Get details of a specific event
// @Tags		Events
// @Produce		json
// @Param		eventId	path		string	true	"Event Id"
// @Success		200		{object}	models.Event
// @Failure		404		{object}	responses.ErrorResponse
// @Router		/events/{eventId} [get]
func (h *EventHandler) FindEvent(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("eventID")
	eventId, err := strconv.ParseInt(pathValue, 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	functionName := fmt.Sprintf("%s.%s",
		config.AppConfig.Schemas["Public"],
		config.AppConfig.DatabaseFunctions["FindEvent"])

	eventMap, err := h.DbService.RunFunction(context.Background(), functionName, map[string]interface{}{"_event_id": eventId})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"event":   eventMap[0],
		"message": "Event retrieved successfully",
		"status":  http.StatusOK,
	})
}

// CreateEvent  @Summary Create event
// @Description	Create a new event
// @Tags		Events
// @Accept		json
// @Produce		json
// @Success		201		{object}	models.Event
// @Router		/events [post]
func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusBadRequest)
		return
	}

	helpers.ValidateRequestData(w, event)

	args := map[string]interface{}{
		"_name":        event.Name,
		"_description": event.Description,
		"_location":    event.Location,
		"_created_at":  event.DateTime,
		"_user_id":     1,
	}

	procedureName := fmt.Sprintf("%s.%s",
		config.AppConfig.Schemas["Public"],
		config.AppConfig.DatabaseProcedures["CreateEvent"])

	err = h.DbService.RunProcedure(context.Background(), procedureName, args)
	if err != nil {
		log.Printf("Failed to create event: %v", err)
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Event created successfully",
		"status":  http.StatusCreated,
	})
}

// UpdateEvent  @Summary Update event
// @Description	Update details of a specific event
// @Tags		Events
// @Accept		json
// @Produce		json
// @Param		eventId	path		string						true	"Event ID"
// @Param		event	body		requests.UpdateEventRequest	true	"Updated event details"
// @Success		200		{object}	models.Event
// @Router		/events/{eventId} [patch]
func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("eventId")
	eventId, err := strconv.ParseInt(pathValue, 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}
	updateEventRequest := requests.UpdateEventRequest{}
	err = json.NewDecoder(r.Body).Decode(&updateEventRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusBadRequest)
		return
	}
	helpers.ValidateRequestData(w, updateEventRequest)
	args := map[string]interface{}{
		"_event_id":    eventId,
		"_name":        updateEventRequest.Name,
		"_description": updateEventRequest.Description,
		"_location":    updateEventRequest.Location,
	}
	procedureName := fmt.Sprintf("%s.%s", config.AppConfig.Schemas["Public"], config.AppConfig.DatabaseProcedures["UpdateEvent"])

	err = h.DbService.RunProcedure(context.Background(), procedureName, args)
	if err != nil {
		log.Printf("Failed to update event: %v", err)
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Event updated successfully",
		"status":  http.StatusOK,
	})
}

// DeleteEvent  @Summary Delete event
// @Description	Deletes an event
// @Tags		Events
// @Produce		json
// @Success		204
// @Param		eventId	path string	true "Event ID"
// @Router		/events/{eventId} [delete]
func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("eventId")
	eventId, err := strconv.ParseInt(pathValue, 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}
	log.Println(eventId)
}
