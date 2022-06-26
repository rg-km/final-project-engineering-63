package helper

import (
	"go_jwt/model/domain"
	"go_jwt/model/web"
)

func ToRegisterResponse(user domain.UserDomain) web.RegisterResponse {
	return web.RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToLoginResponse(user domain.UserDomain, token string) web.LoginResponse {
	return web.LoginResponse{
		Role:  user.Role,
		Token: token,
	}
}

func ToQuizResponse(quiz domain.QuizDomain, category domain.CategoryDomain, answers web.QuizRequest) web.QuizResponse {
	return web.QuizResponse{
		Id:       quiz.Id,
		Category: category.Category,
		Question: quiz.Question,
		Answer:   ToAnswerResponses(answers),
	}
}

func ToAnswerResponses(answers web.QuizRequest) web.AnswerResponse {
	return web.AnswerResponse{
		AnswerA: answers.AnswerA,
		AnswerB: answers.AnswerB,
		AnswerC: answers.AnswerC,
	}
}

func ToQuizResponses(quizRespResult []web.QuizResponse) []web.QuizResponse {
	var quizResponses []web.QuizResponse
	for _, quiz := range quizRespResult {
		quizResponses = append(quizResponses, web.QuizResponse{
			Id:       quiz.Id,
			Question: quiz.Question,
			Category: quiz.Category,
			Answer:   quiz.Answer,
		})
	}
	return quizResponses
}
