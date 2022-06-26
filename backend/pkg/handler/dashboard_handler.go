package handler

import (
	"encoding/json"
	"go_jwt/helper"
	"go_jwt/model/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) createQuiz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var questionReq web.QuizRequest
	err := json.NewDecoder(r.Body).Decode(&questionReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	err = questionReq.ValidateQuiz()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	quizResponse, err := h.services.DashboardService.CreateQuiz(questionReq)
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
		Data:    quizResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (h *Handler) getAdminQuestions(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	quizResponse, err := h.services.DashboardService.GetAdminQuestions()
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
		Data:    quizResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (h *Handler) deleteAdminQuestionById(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	questionIdParam := ps.ByName("questionId")
	questionId, _ := strconv.Atoi(questionIdParam)
	if questionIdParam == "" {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    "questionId is required",
		})
		return
	}

	_, err := h.services.DashboardService.DeleteAdminQuestionById(
		int32(questionId),
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
		Message: "DELETE QUESTION SUCCESSFULLY",
		Data:    nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
