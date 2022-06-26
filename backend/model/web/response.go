package web

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Role  string `json:"role"`
	Token string `json:"token"`
}

type QuizResponse struct {
	Id       int32          `json:"id"`
	Question string         `json:"question"`
	Category string         `json:"category"`
	Answer   AnswerResponse `json:"answer"`
}

type AnswerResponse struct {
	AnswerA string `json:"a"`
	AnswerB string `json:"b"`
	AnswerC string `json:"c"`
}
