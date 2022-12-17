package validation

type Service interface {
	IsValidPassword(pwd string, verification string) bool
}

type service struct {
	factory factory
}

func New(f factory) Service {
	return &service{
		factory: f,
	}
}

func (s service) IsValidPassword(pwd string, verification string) bool {
	return s.factory.Select(verification).Validate(password(pwd))
}
