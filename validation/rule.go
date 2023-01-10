package validation

//go:generate mockery --case=underscore --inpackage --name=Rule
type Rule interface {
	Run(password Password) (invalidBecause string)
}

type rule func(password Password) (invalidBecause string)

func (r rule) Run(password Password) (invalidBecause string) {

	return r(password)
}
