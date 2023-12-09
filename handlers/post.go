package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/santosh1608/project-rss/dataConnector"
	"github.com/santosh1608/project-rss/responder"
)

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	feedId := chi.URLParam(r, "feedId")
	response, err := dataConnector.GetAllPostsByFeedId(feedId)
	if err != nil {
		responder.RespondWithError(w, 400, err.Error())
		return
	}
	responder.RespondWithJSON(w, 200, response)
}

func FetchPostByFeedId(w http.ResponseWriter, r *http.Request) {
	feedId := chi.URLParam(r, "feedId")
	postId := chi.URLParam(r, "postId")
	response, err := dataConnector.GetPostByFeedId(feedId, postId)
	if err != nil {
		responder.RespondWithError(w, 400, err.Error())
		return
	}
	responder.RespondWithJSON(w, 200, response)
}
