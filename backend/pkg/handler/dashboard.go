package handler

type DashboardErrorResponse struct {
	Error string `json:"error"`
}

type QuizItem struct {
	QuizName 	string `json:"quiz_name"`
	Category    string `json:"category"`
	Score       int    `json:"score"`
	Quantity    int    `json:"quantity"`
}

type DashboardSuccessResponse struct {
	Username     string     `json:"username"`
	QuizItems    []QuizItem `json:"quiz_items"`
	TotalScore   int        `json:"total_score"`
	ScoreChanges int        `json:"score_changes"`
}
