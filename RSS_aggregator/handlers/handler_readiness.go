package handlers

import "net/http"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, 200, struct{}{})
}
