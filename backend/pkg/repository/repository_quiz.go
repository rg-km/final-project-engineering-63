package repository

import "database/sql"

type QuizRepository struct {
	db *sql.DB
}

func NewQuizRepository(db *sql.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (p *QuizRepository) FetchQuizByID(id int64) (Quiz, error) {
	sqlStatement := `SELECT id, quiz_name, category FROM quizzes WHERE id = ?`

	row := p.db.QueryRow(sqlStatement, id)

	var quiz Quiz

	err := row.Scan(
		&quiz.ID,
		&quiz.QuizName,
		&quiz.Category,
	)

	if err != nil {
		return Quiz{}, err
	}
	return quiz, nil
}

func (p *QuizRepository) FetchQuizByName(quizName string) (Quiz, error) {
	sqlStatement := `SELECT id, quiz_name, category FROM quizzes WHERE product_name = ?`

	row := p.db.QueryRow(sqlStatement, quizName)

	var quiz Quiz

	err := row.Scan(
		&quiz.ID,
		&quiz.QuizName,
		&quiz.Category,
	)

	if err != nil {
		return Quiz{}, err
	}
	return quiz, nil
}

func (p *QuizRepository) FetchQuizzes() ([]Quiz, error) {
	var sqlStatement string
	var quizzes []Quiz


	sqlStatement = `SELECT id, quiz_name, category FROM quizzes`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	for rows.Next() {
		var quiz Quiz

		err := rows.Scan(
			&quiz.ID,
			&quiz.QuizName,
			&quiz.Category,
		)
		if err != nil {
			return nil, err
		}
	}

	return quizzes, nil
}
