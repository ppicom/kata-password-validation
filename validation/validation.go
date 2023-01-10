package validation

type Service interface {
	IsValidPassword(pwd string) bool
}

type service struct {
	factory factory
}

func New(f factory) Service {
	return &service{
		factory: f,
	}
}

func (s service) IsValidPassword(pwd string) bool {
	allValidators := s.factory.All()
	i := 0
	var isValid bool

	for ; !isValid && i < len(allValidators); i++ {
		isValid = allValidators[i].Validate(password(pwd))
	}

	return isValid
}
