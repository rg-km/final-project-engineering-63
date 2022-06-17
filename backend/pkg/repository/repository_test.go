package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-63/backend/model/domain"
	"github.com/rg-km/final-project-engineering-63/backend/pkg/repository"
)

var _ = Describe("Repository Test", func() {
	var db *sql.DB
	var err error
	var authRepo *repository.AuthRepositorySQLite

	BeforeEach(func() {
		db, err = sql.Open("sqlite3", "./repository_test.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			CREATE TABLE users (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				username TEXT,
				email TEXT,
				password TEXT,
				phone TEXT,
				role TEXT,
				is_login BOOLEAN,
				created_at DATETIME,
				updated_at DATETIME
			);

			INSERT INTO users (
				username, email, password, phone, role, is_login, created_at, updated_at
			) VALUES 
			('admin', 'admin@gmail.com', 'admin123', 'admin', 0, '2022-06-13 00:00:00', '2022-06-04 00:00:00'),
			('rudi', 'rudi@gmail.com', '1234', 'guest', 0, '2022-06-13 00:00:00', '2022-06-04 00:00:00');
		`)

		if err != nil {
			panic(err)
		}

		authRepo = repository.NewAuthRepository(db)

	})

	AfterEach(func() {
		db, err = sql.Open("sqlite3", "./repository_test.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			DROP TABLE users;
		`)

		if err != nil {
			panic(err)
		}
	})

	Describe("GetUser", func() {
		When("email is valid", func() {
			It("should return user", func() {
				user, err := authRepo.GetUser(context.TODO(), "admin@gmail.com", "admin123")
				Expect(err).ToNot(HaveOccurred())

				fmt.Println("user: ", user)
				Expect(user.Username).To(Equal("admin"))
				Expect(user.Email).To(Equal("admin@gmail.com"))
				Expect(user.Password).To(Equal("admin123"))
				Expect(user.Role).To(Equal("admin"))
				Expect(user.IsLogin).To(Equal(false))
			})
		})

		When("email is invalid", func() {
			It("should return error", func() {
				_, err := authRepo.GetUser(context.TODO(), "adminhhua@gmail.com", "1234")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("login failed"))
			})
		})
	})

	Describe("Save", func() {
		When("Save user successfully", func() {
			It("should return user", func() {
				user, err := authRepo.Save(context.TODO(), domain.UserDomain{
					Id:        2,
					Username:  "rudi",
					Email:     "rudi@gmail.com",
					Password:  "1234",
					Role:      "guest",
					IsLogin:   false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
				Expect(err).ToNot(HaveOccurred())

				Expect(user.Username).To(Equal("rudi"))
				Expect(user.Email).To(Equal("rudi@gmail.com"))
				Expect(user.Password).To(Equal("1234"))
				Expect(user.Role).To(Equal("guest"))
				Expect(user.IsLogin).To(Equal(false))
			})
		})
	})

	Describe("Logout", func() {
		When("Logout user successfully", func() {
			It("Accept the logout", func() {
				user, err := authRepo.GetUser(context.TODO(), "admin@gmail.com", "admin123")
				Expect(err).ToNot(HaveOccurred())
				Expect(user.IsLogin).To(Equal(false))

				userIdLogin := user.Id
				isSuccesLogout, err := authRepo.Logout(context.TODO(), userIdLogin)
				Expect(err).ToNot(HaveOccurred())
				Expect(isSuccesLogout).To(Equal(true))
			})
		})
	})
})
