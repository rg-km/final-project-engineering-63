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
	signingKey   = "jwt-secret-key"
	TokenExpires = time.Minute * 30
)

type AuthServiceSQLite struct {
	repo repository.AuthRepository
}

type Claims struct {
	jwt.StandardClaims
	Role   string `json:"role"`
	UserId int    `json:"user_id"`
}

func NewAuthService(repo repository.AuthRepository) *AuthServiceSQLite {
	return &AuthServiceSQLite{
		repo: repo,
	}
}

func (service *AuthServiceSQLite) Create(request web.RegisterCreateRequest, email string) (web.RegisterResponse, error) {
	hashed := security.GeneratePasswordHash(request.Password)

	user := domain.UserDomain{
		Username:  request.Username,
		Email:     "",
		Password:  hashed,
		Role:      "guest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userResult, err := service.repo.Save(user, email)
	helper.PanicIfError(err)

	return helper.ToRegisterResponse(userResult), nil
}

func (service *AuthServiceSQLite) GenerateToken(email, password string) (web.LoginResponse, error) {
	user, err := service.repo.FindUser(
		email, security.GeneratePasswordHash(password),
	)
	helper.PanicIfError(err)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpires).Unix(),
			Issuer:    "APLLICATION",
			IssuedAt:  time.Now().Unix(),
		},
		UserId: int(user.Id),
		Role:   user.Role,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	helper.PanicIfError(err)

	return helper.ToLoginResponse(user, tokenString), nil
}

func (service *AuthServiceSQLite) ParseToken(ctx context.Context, token string) (int32, string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok || !tkn.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	newCtx := context.WithValue(ctx, "user_id", claims.UserId)
	ctx = context.WithValue(ctx, "role", claims.Role)
	ctx = context.WithValue(newCtx, "props", claims)

	return int32(claims.UserId), claims.Role, nil
}
