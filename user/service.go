package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUserInput(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	UserRolesCreate(input UserRolesInput) (UserRoles, error)
	GetAllListUser() ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUserInput(input RegisterUserInput) (User, error) {
	//initiate user struct
	user := User{}
	user.FullName = input.FullName
	user.NIK = input.NIK
	user.CoreUser = input.CoreUser
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (s *service) Login(input LoginInput) (User, error) {
	nik := input.NIK
	password := input.Password

	user, err := s.repository.FindByNIK(nik)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UserRolesCreate(input UserRolesInput) (UserRoles, error) {
	userRoles := UserRoles{}
	userRoles.UserID = input.UserID
	// userRoles.ID = input.ID
	userRoles.RoleCode = input.RoleCode

	userRoles, err := s.repository.SaveUserRoles(userRoles)

	if err != nil {
		return userRoles, err
	}
	return userRoles, nil
}

func (s *service) GetAllListUser() ([]User, error) {

	listUser, err := s.repository.GetAllListUser()
	if err != nil {
		return listUser, err
	}
	return listUser, nil
}
