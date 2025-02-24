package requests

// swagger:model
type FindEventRequest struct {
	EventID string `json:"event_id" validate:"required,sqids"`
}
