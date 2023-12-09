package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/santosh1608/project-rss/dataConnector"
	"github.com/santosh1608/project-rss/responder"
)

func FollowFeed(w http.ResponseWriter, r *http.Request) {
	feedId := chi.URLParam(r, "feedId")
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := claims["userId"].(string)
	_, err := dataConnector.FollowFeed(userId, feedId)
	if err != nil {
		responder.RespondWithError(w, 400, err.Error())
		return
	}
	responder.RespondWithJSON(w, 200, "followed successfully")
}
