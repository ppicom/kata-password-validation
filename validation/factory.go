package validation

//go:generate mockery --exported --name=factory
type factory interface {
	With(validators []Validator) factory
	Select(name string) Validator
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
