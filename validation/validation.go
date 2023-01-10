package validation

type Service interface {
	IsValidPassword(pwd string) (bool, []string)
}

type service struct {
	factory factory
}

func New(f factory) Service {
	return &service{
		factory: f,
	}
}

func (s service) IsValidPassword(pwd string) (bool, []string) {
	allValidators := s.factory.All()
	i := 0
	var (
		isValid        bool
		invalidBecause []string = make([]string, 0)
		reasons        []string
	)

	for ; !isValid && i < len(allValidators); i++ {
		isValid, reasons = allValidators[i].Validate(password(pwd))
		if len(reasons) > 0 {
			invalidBecause = append(invalidBecause, reasons...)
		}
	}

	return isValid, invalidBecause
}
