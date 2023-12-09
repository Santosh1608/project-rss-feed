package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/santosh1608/project-rss/dataConnector"
	dto "github.com/santosh1608/project-rss/dtos"
	"github.com/santosh1608/project-rss/helpers/auth"
	"github.com/santosh1608/project-rss/requests"
	"github.com/santosh1608/project-rss/responder"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var request *requests.Login
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responder.RespondWithError(w, 400, "Error decoding the request")
	}

	user, err := dataConnector.GetUserByEmail(request.Email)

	if err != nil {
		responder.RespondWithError(w, 400, err.Error())
		return
	}

	if !auth.ComparePasswords(user.Password, request.Password) {
		responder.RespondWithError(w, 400, "User not found")
		return
	}

	token, err := auth.CreateToken(user.Id)

	if err != nil {
		responder.RespondWithError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	if err != nil {
		responder.RespondWithError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	response := dto.User{Id: user.Id, Email: user.Email, Name: user.Name, Token: token}
	responder.RespondWithJSON(w, 200, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var request *requests.Register
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responder.RespondWithError(w, 400, "Error decoding the request")
		return
	}

	request.Password = auth.HashPassword(request.Password)

	user, err := dataConnector.Register(request)

	if err != nil {
		responder.RespondWithError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	token, err := auth.CreateToken(user.Id)

	if err != nil {
		responder.RespondWithError(w, 400, fmt.Sprintf("Error: %v", err))
		return
	}

	response := dto.User{Id: user.Id, Email: user.Email, Name: user.Name, Token: token}
	responder.RespondWithJSON(w, 200, response)
}
