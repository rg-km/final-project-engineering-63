package handler

import (
	"encoding/json"
	"go_jwt/pkg/repository"
	"net/http"

	"github.com/rg-km/final-project-engineering-63/backend/pkg/repository"
)

type QuizErrorResponse struct {
	Error string `json:"error"`
}

type AddToQuizSuccessResponse struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type AddToQuizRequest struct {
	QuizName string `json:"quiz_name"`
}

type QuizListSuccessResponse struct {
	QuizItems []repository.QuizItem `json:"quiz_items"`
}

type ScoreRequest struct {
	Score int `json:"score"`
}

type QuizSuccessResponse struct {
	Score    int `json:"score"`
	Total   int `json:"total_payment"`
}

func (handler *Handler) addToQuiz(w http.ResponseWriter, req *http.Request) {
	handler.AllowOrigin(w, req)

	var requestBody AddToQuizRequest
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)

	quiz, err := handler.quizRepo.FetchQuizByName(requestBody.QuizName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(QuizErrorResponse{Error: err.Error()})
		return
	}

	quizItem, err := handler.quizItemRepo.FetchCartByProductID(quiz.ID)
	if err == nil && quizItem.ID != 0 {
		err = handler.quizItemRepo.IncrementCartItemQuantity(quizItem)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(QuizErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(AddToQuizSuccessResponse{
			Name:     quiz.QuizName,
			Category: quiz.Category,
		})
		return
	}

	err = handler.quizItemRepo.InsertQuizItem(repository.QuizItem{
		QuizID: quiz.ID,
		Quantity:  1,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(QuizErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddToQuizSuccessResponse{
		Name:     quiz.QuizName,
		Category: quiz.Category,
	})
}

func (handler *Handler) clearQuiz(w http.ResponseWriter, req *http.Request) {
	handler.AllowOrigin(w, req)
	err := handler.quizItemRepo.ResetQuizItems()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(QuizErrorResponse{Error: err.Error()})
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) quizList(w http.ResponseWriter, req *http.Request) {
	handler.AllowOrigin(w, req)
	quizItems, err := handler.quizItemRepo.FetchQuizItems()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(QuizErrorResponse{Error: err.Error()})
		}
	}()

	encoder.Encode(QuizListSuccessResponse{QuizItems: quizItems})
}

func (handler *Handler) score(w http.ResponseWriter, req *http.Request) {
	handler.AllowOrigin(w, req)

	var requestBody ScoreRequest

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	quizItems, err := handler.quizItemRepo.FetchQuizItems()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(QuizErrorResponse{Error: err.Error()})
		}
	}()

	sumScore, err := handler.quizItemRepo.TotalScore()
	if err != nil {
		return
	}

	scoreChanges, err := handler.quizItemRepo.ansactionRepo.Pay(quizItems, requestBody.Score)
	if err != nil {
		return
	}

	encoder.Encode(QuizSuccessResponse{
		Score: scoreChanges,
		Total:   sumScore,
		
	})
}
