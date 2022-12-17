package validation

//go:generate mockery --case=underscore --inpackage --name=Rule
type Rule interface {
	Run(password Password) bool
}

type rule func(password Password) bool

func (r rule) Run(password Password) bool {

	return r(password)
}
