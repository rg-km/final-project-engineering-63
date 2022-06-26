package repository

import (
	"database/sql"
	"go_jwt/helper"
	"go_jwt/model/domain"
	"go_jwt/model/web"
)

type HomeRepositorySQLite struct {
	DB *sql.DB
}

func NewHomeRepository(db *sql.DB) *HomeRepositorySQLite {
	return &HomeRepositorySQLite{
		DB: db,
	}
}

func (repo *HomeRepositorySQLite) FindCategoryByName(name string) (domain.CategoryDomain, error) {
	var category domain.CategoryDomain

	query := `SELECT id, category FROM categories WHERE category = ?`
	row := repo.DB.QueryRow(query, name)
	err := row.Scan(&category.Id, &category.Category)
	helper.PanicIfError(err)

	return category, nil
}

func (repo *HomeRepositorySQLite) FindCategoryById(categoryId int32) (domain.CategoryDomain, error) {
	var category domain.CategoryDomain

	query := `SELECT id, category FROM categories WHERE id = ?`
	row := repo.DB.QueryRow(query, categoryId)
	err := row.Scan(&category.Id, &category.Category)
	helper.PanicIfError(err)

	return category, nil
}

func (repo *HomeRepositorySQLite) FindQuizByCategoryIdWithPagination(category domain.CategoryDomain, page, limit int32) ([]web.QuizResponse, error) {
	var quiz []web.QuizResponse

	query := `
		SELECT q.id, c.category, q.question, a.answer_a, a.answer_b, a.answer_c
		FROM quizzes AS q
		INNER JOIN categories AS c ON c.id = q.category_id
		INNER JOIN answers AS a ON q.id = a.quiz_id
		WHERE category_id = ?
		ORDER BY q.id
		LIMIT ? OFFSET ?;`

	rows, err := repo.DB.Query(query, category.Id, limit, (page-1)*limit)
	helper.PanicIfError(err)

	for rows.Next() {
		var quizResponse web.QuizResponse

		err := rows.Scan(
			&quizResponse.Id, &quizResponse.Category, &quizResponse.Question,
			&quizResponse.Answer.AnswerA, &quizResponse.Answer.AnswerB,
			&quizResponse.Answer.AnswerC,
		)
		helper.PanicIfError(err)

		quiz = append(quiz, quizResponse)
	}

	if err := rows.Err(); err != nil {
		helper.PanicIfError(err)
	}

	return quiz, nil
}
