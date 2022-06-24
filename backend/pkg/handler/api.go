package handler

import (
	"fmt"
	"go_jwt/pkg/repository"
	"net/http"
)

type API struct {
	usersRepo    repository.UserRepository
	Repo         repository.Repository
	quizRepo     repository.QuizRepository
	quizItemRepo repository.QuizItemRepository
	mux          *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, Repo repository.Repository, quizRepo repository.QuizRepository, quizItemRepo repository.QuizItemRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, Repo, quizRepo, quizItemRepo, mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))

	// API with AuthMiddleware:
	mux.Handle("/api/quiz", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.quizList))))
	mux.Handle("/api/quiz/add", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.addToQuiz))))
	mux.Handle("/api/quiz/clear", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.clearQuiz))))
	mux.Handle("/api/quizItem", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.quizList))))
	mux.Handle("/api/start", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.start))))

	// API with AuthMiddleware and AdminMiddleware
	mux.Handle("/api/admin/quiz", api.GET(api.AuthMiddleWare(api.AdminMiddleware(http.HandlerFunc(api.getDashboard)))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
