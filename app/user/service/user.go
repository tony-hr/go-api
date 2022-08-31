package service

import (
	"errors"
	"go-api/app/user/repository"
	"go-api/middleware"
	"go-api/requests"
	resp "go-api/responses"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(*requests.LoginRequest) (*resp.LoginResp, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) Login(req *requests.LoginRequest) (*resp.LoginResp, error) {
	users, err := u.repo.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	//check password
	checkPassword := checkPasswordHash(req.Password, users.Password)
	if checkPassword == false {
		return nil, errors.New("Wrong username or password")
	}

	token, err := middleware.GenerateToken(req.Username)
	if err != nil {
		return nil, err
	}

	return &resp.LoginResp{
		Username:    req.Username,
		Fullname:    users.Fullname,
		Email:       users.Email,
		PhoneNumber: users.PhoneNumber,
		Token:       token,
	}, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPasswordUser(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}
