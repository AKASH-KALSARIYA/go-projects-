package service

import (
	"auth/model"
	"auth/repository"
	"auth/utill"
	"errors"
)

func RegisterUser(user model.User) error {
	hashed, err := utill.HashPassword(user.PassWord)
	if err != nil {
		return err
	}

	user.PassWord = hashed
	return repository.InsertUser(user)
}

func LoginUser(email, password string) (string, error) {
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utill.CheckPasswordHash(password, user.PassWord) {
		return "", errors.New("invalid password")
	}

	token, err := utill.GenerateToken(user.Email)
	return token, err
}
