package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go-net_http-project/internal/handlers"
	"net/http"
)

func RegisterRoutes(router *http.ServeMux, dbPool *pgxpool.Pool) {
	commonHandlers := handlers.NewCommonHandlers(dbPool)
	eventHandler := handlers.NewEventHandler(commonHandlers)

	router.HandleFunc("GET /events", eventHandler.GetEvents)
	router.HandleFunc("GET /events/{eventId}", eventHandler.FindEvent)
	router.HandleFunc("POST /events", eventHandler.CreateEvent)
	router.HandleFunc("PATCH /events/{eventId}", eventHandler.UpdateEvent)
	router.HandleFunc("DELETE /events/{eventId}", eventHandler.DeleteEvent)
}
