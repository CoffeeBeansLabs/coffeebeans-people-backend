package router

import (
	"coffeebeans-people-backend/dao"
	"github.com/go-chi/chi"
	"net/http"
	"github.com/rs/cors"
)

// API is an API configuration.
type API struct {
	DaoService *dao.Service
}

// APIMux returns an API multiplexer.
func APIMux(api *API) *chi.Mux {
	mux := chi.NewMux()
	corsSettings := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"X-XSRF-Token",
			"X-HTTP-Method-Override",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           86400,
	})

	mux.Use(corsSettings.Handler)

	mux.Route("/ping", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("pong"))
		})
	})

	return mux

}
