package feedback

type UseCase interface {
	Store(f Feedback) error
}

type Service struct {}

func NewService() *Service {
	return &Service{}
}
func (s *Service) Store(f Feedback) error {
	//@TODO create store rules, using databases or something else
	return nil
}
