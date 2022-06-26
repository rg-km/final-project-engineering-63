package domain

import "time"

type UserDomain struct {
	Id        int32     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryDomain struct {
	Id        int32     `json:"id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QuizDomain struct {
	Id         int32     `json:"id"`
	CategoryId int32     `json:"category_id"`
	Question   string    `json:"question"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AnswerDomain struct {
	Id         int32     `json:"id"`
	UserId     int32     `json:"user_id"`
	QuizId     int32     `json:"quiz_id"`
	AnswerA    string    `json:"answer_a"`
	AnswerB    string    `json:"answer_b"`
	AnswerC    string    `json:"answer_c"`
	AnswerTrue string    `json:"answer_true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AnswerAttempt struct {
	UserId int32  `json:"user_id"`
	Answer string `json:"answer"`
}
