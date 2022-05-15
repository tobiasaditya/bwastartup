package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(inputLogin LoginInput) (User, error)
	CheckEmailUser(input EmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}

}

func (s service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	createdUser, err := s.repository.Save(user)
	if err != nil {
		return createdUser, err
	}

	return createdUser, nil
}

func (s service) LoginUser(inputLogin LoginInput) (User, error) {
	email := inputLogin.Email
	password := inputLogin.Password

	foundUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return foundUser, err
	}

	// Check if user found or not
	if foundUser.ID == 0 {
		return foundUser, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password))
	if err != nil {
		return foundUser, err
	}

	return foundUser, nil
}

func (s service) CheckEmailUser(input EmailInput) (bool, error) {
	email := input.Email

	foundUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// Check if user found or not
	if foundUser.ID != 0 {
		return false, nil
	}

	return true, nil
}

func (s service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// Found user by id
	foundUser, err := s.repository.FindByID(ID)
	if err != nil {
		return foundUser, err
	}

	// Check if user found or not
	if foundUser.ID == 0 {
		return foundUser, errors.New("User not found")
	}

	foundUser.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(foundUser)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s service) GetUserByID(ID int) (User, error) {
	foundUser, err := s.repository.FindByID(ID)
	if err != nil {
		return foundUser, err
	}
	return foundUser, err
}
