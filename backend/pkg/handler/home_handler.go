package handler

import (
	"encoding/json"
	"go_jwt/helper"
	"go_jwt/model/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) getQuizByCategoryIdWithPagination(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categoryIdParam := ps.ByName("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdParam)
	if categoryIdParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    "categoryId is required",
		})
		return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	quizResponses, err := h.services.HomeService.GetQuizByCategoryIdWithPagination(
		int32(categoryId), int32(page), int32(limit),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    quizResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
