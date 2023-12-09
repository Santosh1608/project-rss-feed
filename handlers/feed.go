package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/santosh1608/project-rss/dataConnector"
	dto "github.com/santosh1608/project-rss/dtos"
	"github.com/santosh1608/project-rss/models"
	"github.com/santosh1608/project-rss/requests"
	"github.com/santosh1608/project-rss/responder"
)

func CreateFeed(w http.ResponseWriter, r *http.Request) {
	var request = requests.CreateFeed{}
	fmt.Println(r.Body, "request body is")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responder.RespondWithError(w, 400, "Error decoding the request")
		return
	}
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := claims["userId"].(string)
	feed, err := dataConnector.CreateFeed(&models.Feed{UserId: userId, Name: request.Name, Url: request.Url})
	if err != nil {
		responder.RespondWithError(w, 400, err.Error())
		return
	}
	response := dto.Feed{Id: feed.Id, Name: feed.Name, UserId: feed.UserId}
	responder.RespondWithJSON(w, 200, response)
}
