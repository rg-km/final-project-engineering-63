package service

import (
	"context"
	"fmt"
	"go_jwt/helper"
	"go_jwt/model/domain"
	"go_jwt/model/web"
	"go_jwt/pkg/repository"
	"go_jwt/security"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt         = "kdjsdsjdfsjfiosidasdkskdfksnd"
	signingKey   = "secret_key"
	TokenExpires = time.Minute * 30
)

type AuthService struct {
	repo repository.UserRepository
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (service *AuthService) Create(ctx context.Context, request web.RegisterCreateRequest) (web.RegisterResponse, error) {
	hashed := security.GeneratePasswordHash(request.Password)

	user := domain.UserDomain{
		Username: request.Username,
		Email:    request.Email,
		Password: hashed,
		Role:     "guest",
		IsLogin:  false,
	}

	userResult, err := service.repo.Save(ctx, user)
	helper.PanicIfError(err)

	return helper.ToRegisterResponse(userResult), nil
}

func (service *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := service.repo.GetUser(context.Background(), email, security.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpires).Unix(),
			Issuer:    "todo-app",
			IssuedAt:  time.Now().Unix(),
		},
		Username: user.Username,
		UserId:   int(user.Id),
	})

	return token.SignedString([]byte(signingKey))
}

func (service *AuthService) ParseToken(token string) (string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok || !tkn.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.Username, nil
}

func (service *AuthService) Logout(ctx context.Context, username string) (bool, error) {
	return service.repo.Logout(ctx, username)
}
