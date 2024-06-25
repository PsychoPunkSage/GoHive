package config

import (
	"RSS_aggregator/handlers"
	"RSS_aggregator/internal/auth"
	"RSS_aggregator/internal/database"
	"fmt"
	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *APIConfig) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			handlers.RespondWithError(w, 403, fmt.Sprintf("Auth Error %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			handlers.RespondWithError(w, 403, fmt.Sprintf("Couldn't get user %v", err))
			return
		}
		handler(w, r, user)
	}
}
