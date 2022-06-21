package web

import (
	"errors"
	"strings"
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

func (l *LoginCreateRequest) PrepareLogin() {
	l.Email = strings.TrimSpace(l.Email)
	l.Password = strings.TrimSpace(l.Password)
}

func (l *LoginCreateRequest) ValidateLogin(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if l.Email == "" {
			return errors.New("required Email")
		}
		if l.Password == "" {
			return errors.New("required Password")
		}
		return nil
	}

	return errors.New("invalid Action")
}

func (r *RegisterCreateRequest) PrepareRegister() {
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)

}

func (r *RegisterCreateRequest) ValidateRegister(action string) error {
	switch strings.ToLower(action) {
	case "create":

		if r.Email == "" {
			return errors.New("required Email")
		}
		if r.Password == "" {
			return errors.New("required Password")
		}

		return nil
	}

	return errors.New("invalid Action")
}
