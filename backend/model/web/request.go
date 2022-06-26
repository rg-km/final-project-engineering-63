package web

import (
	"errors"
)

type LoginCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type QuizRequest struct {
	Question   string `json:"question"`
	Category   string `json:"category"`
	AnswerA    string `json:"a"`
	AnswerB    string `json:"b"`
	AnswerC    string `json:"c"`
	AnswerTrue string `json:"answer_true"`
}

func (q *QuizRequest) ValidateQuiz() error {
	if q.Question == "" {
		return errors.New("required Question")
	}
	if q.Category == "" {
		return errors.New("required Category")
	}
	if q.AnswerA == "" {
		return errors.New("required AnswerA")
	}
	if q.AnswerB == "" {
		return errors.New("required AnswerB")
	}
	if q.AnswerC == "" {
		return errors.New("required AnswerC")
	}
	if q.AnswerTrue == "" {
		return errors.New("required AnswerTrue")
	}

	return nil
}

func (l *LoginCreateRequest) ValidateLogin() error {
	if l.Email == "" {
		return errors.New("required Email")
	}
	if l.Password == "" {
		return errors.New("required Password")
	}

	return nil
}

func (r *RegisterCreateRequest) ValidateRegister() error {
	if r.Username == "" {
		return errors.New("required Username")
	}
	if r.Email == "" {
		return errors.New("required Email")
	}
	if r.Password == "" {
		return errors.New("required Password")
	}

	return nil
}
