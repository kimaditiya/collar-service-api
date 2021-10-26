package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUserInput(input RegisterUserInput) (User, error)
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
