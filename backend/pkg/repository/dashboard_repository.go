package repository

import (
	"database/sql"
	"go_jwt/helper"
	"go_jwt/model/domain"
	"go_jwt/model/web"
)

type DashboardRepositorySQLite struct {
	DB *sql.DB
}

func NewDashboardRepository(db *sql.DB) *DashboardRepositorySQLite {
	return &DashboardRepositorySQLite{
		DB: db,
	}
}

func (repo *DashboardRepositorySQLite) SaveQuiz(quizDomain domain.QuizDomain, quizReq web.QuizRequest) (domain.QuizDomain, error) {
	query := `
		INSERT INTO quizzes (category_id, question, answer_true) 
		VALUES (?, ?, ?);`

	result, err := repo.DB.Exec(query,
		quizDomain.CategoryId, quizDomain.Question, quizReq.AnswerTrue,
	)
	helper.PanicIfError(err)

	quizId, err := result.LastInsertId()
	helper.PanicIfError(err)

	quizDomain.Id = int32(quizId)

	query = `
		INSERT INTO answers 
		(quiz_id, answer_a, answer_b, answer_c) 
		VALUES (?, ?, ?, ?);`

	result, err = repo.DB.Exec(query,
		quizDomain.Id, quizReq.AnswerA, quizReq.AnswerB, quizReq.AnswerC,
	)
	helper.PanicIfError(err)

	_, err = result.LastInsertId()
	helper.PanicIfError(err)

	return quizDomain, nil
}

func (repo *DashboardRepositorySQLite) FindAllAdminQuestions() ([]web.QuizResponse, error) {
	query := `
		SELECT q.id, q.question, c.category, a.answer_a, a.answer_b, a.answer_c
		FROM quizzes AS q
		INNER JOIN categories AS c ON q.category_id = c.id
		INNER JOIN answers AS a ON q.id = a.quiz_id;`

	rows, err := repo.DB.Query(query)
	helper.PanicIfError(err)

	var quizResponses []web.QuizResponse

	for rows.Next() {
		var quizResponse web.QuizResponse

		err := rows.Scan(
			&quizResponse.Id, &quizResponse.Question, &quizResponse.Category,
			&quizResponse.Answer.AnswerA, &quizResponse.Answer.AnswerB,
			&quizResponse.Answer.AnswerC,
		)
		helper.PanicIfError(err)

		quizResponses = append(quizResponses, quizResponse)
	}

	if err := rows.Err(); err != nil {
		helper.PanicIfError(err)
	}

	return quizResponses, nil
}

func (repo *DashboardRepositorySQLite) DeleteAdminQuestionById(questionId int32) (bool, error) {
	query := `DELETE FROM quizzes WHERE id = ?;`

	_, err := repo.DB.Exec(query, questionId)
	helper.PanicIfError(err)

	query = `DELETE FROM answers WHERE quiz_id = ?;`
	_, err = repo.DB.Exec(query, questionId)
	helper.PanicIfError(err)

	return true, nil
}
