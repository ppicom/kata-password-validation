package validation

type Validator interface {
	Validate(password Password) bool
	With(rules RuleSet) Validator
	Equals(validator Validator) bool
	Name() string
}

type validatorPrototype struct {
	rules RuleSet
	name  string
}

func NewValidator(name string) Validator {

	return &validatorPrototype{
		name:  name,
		rules: NewRuleset(),
	}
}

func (v *validatorPrototype) With(rules RuleSet) Validator {

	v.rules = rules
	return v
}

func (v *validatorPrototype) Validate(password Password) bool {

	return v.rules.RunAgainst(password)
}

func (v *validatorPrototype) Equals(validator Validator) bool {

	return v.Name() == validator.Name()
}

func (v *validatorPrototype) Name() string {

	return v.name
}
