package services

import (
	"errors"
	"gin-user-management/lib"
	"gin-user-management/models"
	"gin-user-management/repository"
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	logger     lib.Logger
	repository repository.UserRepository
	env        lib.Env
}

func NewAuthService(logger lib.Logger, repository repository.UserRepository, env lib.Env) AuthService {
	return AuthService{
		logger:     logger,
		repository: repository,
		env:        env,
	}
}

func (authService AuthService) Login(username string) (token string, user models.User, error error) {
	user, err := authService.ValidateUser(username)
	if err != nil {
		return "", models.User{}, err
	}
	token, errs := authService.CreateToken(user)
	if errs != nil {
		return "", models.User{}, err
	}
	return token, user, nil
}

func (authService AuthService) ValidateUser(username string) (user models.User, error error) {
	//todo: authenticate user with password
	return user, authService.repository.First(&user, "username = ?", username).Error
}

func (authService AuthService) CreateNewUser(user models.User) (createdUser models.User, error error) {
	return user, authService.repository.Create(&user).Error
}

func (authService AuthService) CreateToken(user models.User) (tokenString string, error error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"username":  user.Username,
		"full_name": user.FullName,
	})

	tokenString, err := token.SignedString([]byte(authService.env.JWTSecret))

	if err != nil {
		authService.logger.Zap.Error("JWT validation failed: ", err)
	}

	return tokenString, err
}

func (authService AuthService) Authorize(tokenString string) (bool, error, float64,) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(authService.env.JWTSecret), nil
	})
	if token.Valid {
		return true, nil, token.Claims.(jwt.MapClaims)["id"].(float64)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed"), 0
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired"), 0
		}
	}
	return false, errors.New("couldn't handle token"), 0
}
