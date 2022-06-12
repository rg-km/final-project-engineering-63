package helper

import (
	"go_jwt/model/domain"
	"go_jwt/model/web"
)

func ToRegisterResponse(user domain.UserDomain) web.RegisterResponse {
	return web.RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToLoginResponse(role, token string) web.LoginResponse {
	return web.LoginResponse{
		Role:  role,
		Token: token,
	}
}
