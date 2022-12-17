package validation

type Validator interface {
	Validate(password string) bool
	With(rules []Rule) Validator
	Equals(validator Validator) bool
	Name() string
}

type validatorPrototype struct {
	rules []Rule
	name  string
}

func NewValidator(name string) Validator {

	return &validatorPrototype{
		name:  name,
		rules: make([]Rule, 0),
	}
}

func (v *validatorPrototype) With(rules []Rule) Validator {

	v.rules = rules
	return v
}

func (v *validatorPrototype) Validate(password string) bool {

	for _, r := range v.rules {

		if !(r(password)) {
			return false
		}
	}

	return true
}

func (v *validatorPrototype) Equals(validator Validator) bool {
	return v.Name() == validator.Name()
}

func (v *validatorPrototype) Name() string {
	return v.name
}
