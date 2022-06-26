package seed

import (
	"database/sql"
	"go_jwt/helper"
	"go_jwt/security"
	"log"
	"time"
)

func Seed(db *sql.DB) {
	query := `
	INSERT INTO users 
	(username, email, password, role, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?);

	INSERT INTO categories
	(category, created_at, updated_at)
	VALUES
	('VOCABULARY', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('GRAMMAR', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('TENSES', '2020-01-01 00:00:00', '2020-01-01 00:00:00');
	`

	hashed := security.GeneratePasswordHash("admin123")
	formatHours := "2006-01-02 15:04:05"

	args := []interface{}{
		"admin",
		"admin@gmail.com",
		hashed,
		"admin",
		time.Now().Format(formatHours),
		time.Now().Format(formatHours),
	}

	_, err := db.Exec(query, args...)
	helper.PanicIfErrorWithMessage("Error when seed with error:", err)

	log.Println("Seeding success")
}
