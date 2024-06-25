package config

import (
	"RSS_aggregator/handlers"
	"RSS_aggregator/internal/database"
	"RSS_aggregator/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type APIConfig struct {
	DB *database.Queries
}

func (apiCfg *APIConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	handlers.RespondWithJson(w, 201, model.DatabaseUserToUser(user))
}

func (apiCfg *APIConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	handlers.RespondWithJson(w, 200, model.DatabaseUserToUser(user))
}

func (apiCfg *APIConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	handlers.RespondWithJson(w, 201, model.DatabaseFeedToFeed(feed))
}

func (apiCfg *APIConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error fetching feeds: %v", err))
		return
	}

	handlers.RespondWithJson(w, 201, model.DatabaseFeedsToFeeds(feeds))
}

func (apiCfg *APIConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	handlers.RespondWithJson(w, 201, model.DatabaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *APIConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiCfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error fetching feed follows: %v", err))
		return
	}
	handlers.RespondWithJson(w, 201, model.DatabaseFeedFollowsToFeedFollows(feed_follows))
}

func (apiCfg *APIConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		handlers.RespondWithError(w, 400, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	handlers.RespondWithJson(w, 200, struct{}{})
}
