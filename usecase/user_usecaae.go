package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/masa720/todo-backend-golang/model"
	"github.com/masa720/todo-backend-golang/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Signup(user *model.User) error
	Signin(mail string, pass string) (string, error)
	// Signout() error
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Signup(user *model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Pass), 10)
	if err != nil {
		return err
	}
	newUser := model.User{Name: user.Name, Mail: user.Mail, Pass: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) Signin(mail string, pass string) (string, error) {
	user, err := uu.ur.FindByEmail(mail)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(pass)); err != nil {
		return "", err
	}

	// JWTトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JWTClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	// シークレットキーで署名（本番環境では環境変数から取得すべき）
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
