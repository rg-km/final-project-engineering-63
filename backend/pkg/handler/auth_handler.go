package handler

import (
	"encoding/json"
	"go_jwt/helper"
	"go_jwt/model/web"
	"go_jwt/pkg/service"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var registerCreateRequest web.RegisterCreateRequest
	err := json.NewDecoder(r.Body).Decode(&registerCreateRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	err = registerCreateRequest.ValidateRegister()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	userResponse, err := h.services.Create(
		registerCreateRequest, registerCreateRequest.Email,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    userResponse,
	}

	json.NewEncoder(w).Encode(webResponse)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var loginRequest web.LoginCreateRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	err = loginRequest.ValidateLogin()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	loginResponse, err := h.services.AuthService.GenerateToken(
		loginRequest.Email, loginRequest.Password,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    loginResponse,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   loginResponse.Token,
		Expires: time.Now().Add(service.TokenExpires),
		Path:    "/",
	})

	helper.WriteToResponseBody(w, webResponse)
}
