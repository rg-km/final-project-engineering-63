package handler

import (
	"encoding/json"
	"go_jwt/model/web"
	"go_jwt/pkg/service"
	"net/http"
	"time"
)

func (h *Handler) signUp(writer http.ResponseWriter, request *http.Request) {
	var registerCreateRequest web.RegisterCreateRequest
	err := json.NewDecoder(request.Body).Decode(&registerCreateRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	registerCreateRequest.PrepareRegister()
	err = registerCreateRequest.ValidateRegister("create")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	userResponse, err := h.services.Create(
		request.Context(), registerCreateRequest,
	)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
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

	json.NewEncoder(writer).Encode(webResponse)
}

func (h *Handler) signIn(writer http.ResponseWriter, request *http.Request) {
	var loginRequest web.LoginCreateRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	loginRequest.PrepareLogin()
	err = loginRequest.ValidateLogin("login")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	token, err := h.services.UserService.GenerateToken(
		loginRequest.Email, loginRequest.Password,
	)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    token,
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(service.TokenExpires),
		Path:    "/",
	})

	json.NewEncoder(writer).Encode(webResponse)
}

func (h *Handler) signOut(writer http.ResponseWriter, request *http.Request) {
	token, err := request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if token.Value == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	username, _ := h.services.UserService.ParseToken(token.Value)

	_, err = h.services.UserService.Logout(request.Context(), username)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    "Successfully logged out",
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		MaxAge:   -1,
	}

	http.SetCookie(writer, &cookie)

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(webResponse)
}
