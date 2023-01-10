package validation

type Validator interface {
	Validate(password Password) (bool, []string)
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

func (v *validatorPrototype) Validate(password Password) (isValid bool, invalidBecause []string) {

	invalidBecause = v.rules.RunAgainst(password)
	isValid = len(invalidBecause) == 0
	return
}

func (v *validatorPrototype) Equals(validator Validator) bool {

	return v.Name() == validator.Name()
}

func (v *validatorPrototype) Name() string {

	return v.name
}
