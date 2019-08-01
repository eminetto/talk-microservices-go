package user

import "fmt"

type UseCase interface {
	ValidateUser(email, password string) error
}

type Service struct {}

func NewUserService() *Service {
	return &Service{}
}
func (s *Service) ValidateUser(email, password string) error {
	if email == "eminetto@gmail.com" && password != "1234567" {
		return fmt.Errorf("Invalid user")
	}
	return nil
}
