package internal

import (
	"net/http"
	"saasteamtest/backend/data"
	"saasteamtest/backend/domain"
	"saasteamtest/backend/internal/handlers"
	"saasteamtest/backend/models"

	"github.com/go-chi/chi"
)

func RouterInitializer() *chi.Mux {
	db := data.DB{PortTable: map[string]models.Port{}}
	portHandler := data.NewPortHandler(db)
	portService := domain.NewPortService(portHandler)

	r := chi.NewRouter()

	r.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	})

	r.Route("/ports", func(r chi.Router) {
		r.Method("PUT", "/", handlers.BaseHandler(handlers.CreatePort(portService)))
		r.Method("POST", "/", handlers.BaseHandler(handlers.UpdatePort(portService)))
	})

	return r
}
