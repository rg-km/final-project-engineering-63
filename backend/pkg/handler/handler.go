package handler

import (
	"fmt"
	"go_jwt/pkg/service"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (handler *Handler) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	fmt.Println("=========================")
	fmt.Println("Running in localhost:8080")
	fmt.Println("=========================")

	router.POST("/auth/register", handler.signUp)

	router.POST("/auth/login", handler.signIn)

	router.POST("/admin/quizzes/create",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.createQuiz,
			),
		),
	)

	router.GET("/admin/quizzes",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.getAdminQuestions,
			),
		),
	)

	router.DELETE("/admin/quizzes/:questionId",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.deleteAdminQuestionById,
			),
		),
	)

	router.GET("/home/start-quiz/:categoryId",
		handler.AuthMiddleWare(
			handler.getQuizByCategoryIdWithPagination,
		),
	)

	return router
}
