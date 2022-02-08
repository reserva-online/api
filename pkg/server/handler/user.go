package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/schedule-api/pkg/user"
)

type UserSaver interface {
	Save(ctx context.Context, data user.User) (int, error)
}

func HandleUserSave(userSaver UserSaver) http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Type     string `json:"type"`
	}
	type response struct {
		Id int `json:"id"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		var body request
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			log.Printf("User save error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusBadRequest, response)
			return
		}
		data := user.User{
			Email:    body.Email,
			Password: body.Password,
			Type:     body.Type,
		}
		result, err := userSaver.Save(req.Context(), data)
		if err != nil {
			log.Printf("User save error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Id: result,
		}
		log.Printf("User save sucess!")
		makeResponse(res, http.StatusOK, response)
	}
}

type UserGetter interface {
	GetById(ctx context.Context, id int) (user.User, error)
}

func HandleUserGetById(userGetter UserGetter) http.HandlerFunc {
	type response struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Type     string `json:"type"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		idParam, ok := getPathParameterFromRequest(req, "id")
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			log.Printf("User get error - id not informed")
			makeResponse(res, http.StatusBadRequest, &errorResponse{Message: "id not informed"})
			return
		}

		result, err := userGetter.GetById(req.Context(), id)
		if err != nil {
			log.Printf("User get error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Id:    result.ID,
			Email: result.Email,
			Type:  result.Type,
		}
		log.Printf("User get sucess!")
		makeResponse(res, http.StatusOK, response)
	}
}

type UserLoginer interface {
	Login(ctx context.Context, loginUser user.LoginUser) (user.LoginResponse, error)
}

func HandleUserLogin(userLoginer UserLoginer) http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		Token string `json:"token"`
		ID    int    `json:"userId"`
		Type  string `json:"userType"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		var body request
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			log.Printf("User login Error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusBadRequest, response)
			return
		}
		data := user.LoginUser{
			Email:    body.Email,
			Password: body.Password,
		}
		result, err := userLoginer.Login(req.Context(), data)
		if err != nil {
			log.Printf("User login Error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Token: result.Token,
			ID:    result.ID,
			Type:  result.Type,
		}
		log.Printf("User login Success!")
		makeResponse(res, http.StatusOK, response)
	}
}
