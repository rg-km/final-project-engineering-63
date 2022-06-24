package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type QuizItemRepository struct {
	db *sql.DB
}

func NewQuizItemRepository(db *sql.DB) *QuizItemRepository {
	return &QuizItemRepository{db: db}
}

func (c *QuizItemRepository) FetchQuizItems() ([]QuizItem, error) {
	var sqlStatement string
	var quizItems []QuizItem

	sqlStatement = `SELECT 
	c.id,
	p.category,
	c.quiz_id,
	p.quiz_name,
	c.quantity
	FROM quiz_items C
	INNER JOIN quizzes p ON c.quiz_id = p.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	for rows.Next() {
		var quizItem QuizItem

		err := rows.Scan(
			&quizItem.ID,
			&quizItem.Category,
			&quizItem.QuizID,
			&quizItem.QuizName,
			&quizItem.Quantity,
		)
		if err != nil {
			return nil, err
		}
		quizItems = append(quizItems, quizItem)
	}

	return quizItems, nil
}

func (c *QuizItemRepository) FetchQuizByQuizID(quizID int64) (QuizItem, error) {
	var quizItem QuizItem
	var sqlStatement string


	sqlStatement = `SELECT 
	c.id,
	p.category,
	c.quiz_id,
	p.quiz_name,
	c.quantity
	FROM quiz_items C
	INNER JOIN quizzes p ON c.quiz_id = p.id
	WHERE c.quiz_id = ?`

	row := c.db.QueryRow(sqlStatement, quizID)
	
	err := row.Scan(
		&quizItem.ID,
		&quizItem.Category,
		&quizItem.QuizID,
		&quizItem.QuizName,
		&quizItem.Quantity,
	)

	if err != nil {
		return QuizItem{}, err
	}

	return quizItem, nil
}

func (c *QuizItemRepository) InsertQuizItem(quizItem QuizItem) error {
	_, err := c.FetchQuizByQuizID(quizItem.QuizID)
	
	if err != nil {
		sqlStmt := `INSERT INTO quiz_items (quiz_id, quantity)
		VALUES (?, ?)`

		_, err := c.db.Exec(sqlStmt, quizItem.QuizID, quizItem.Quantity)
		if err != nil {
			return err
		}
		return nil
	}

	err = c.IncrementQuizItemQuantity(quizItem)
	if err != nil {
		return err
	}

	return nil
}

func (c *QuizItemRepository) IncrementQuizItemQuantity(quizItem QuizItem) error {
	sqlStmt := `UPDATE quiz_items SET quantity = quantity + ? WHERE quiz_id = ?`

	_, err := c.db.Exec(sqlStmt, quizItem.Quantity, quizItem.QuizID)

	if err != nil {
		return err
	}
	return nil
}

func (c *QuizItemRepository) ResetQuizItems() error {
	sqlStmt := `DELETE FROM quiz_items`

	_, err := c.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}


