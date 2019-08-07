package vote

type UseCase interface {
	Store(v Vote) error
}

type Service struct {}

func NewService() *Service {
	return &Service{}
}
func (s *Service) Store(v Vote) error {
	//@TODO create store rules, using databases or something else
	return nil
}
