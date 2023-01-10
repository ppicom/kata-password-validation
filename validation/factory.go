package validation

//go:generate mockery --case=underscore --inpackage --name=factory
type factory interface {
	With(validators []Validator) factory
	Select(name string) Validator
	All() []Validator
}

type impl struct {
	validators map[string]Validator
}

func Factory() factory {

	return &impl{
		validators: make(map[string]Validator, 0),
	}
}

func (f *impl) With(validators []Validator) factory {

	for _, v := range validators {
		f.validators[v.Name()] = v
	}
	return f
}

func (f *impl) Select(name string) Validator {

	return f.validators[name]
}

// All implements factory
func (f *impl) All() []Validator {
	all := make([]Validator, 0)

	for _, v := range f.validators {
		all = append(all, v)
	}

	return all
}
