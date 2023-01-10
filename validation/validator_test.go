package validation

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidatorTestSuite struct {
	suite.Suite
}

func (suite *ValidatorTestSuite) TestValidatorShouldEnforceARuleOnAPassword() {

	password := password("pass")
	noRulesValidator := NewValidator("empty_validator")
	lengthValidator := NewValidator("length_validator").With(NewRuleset().ForLength(5))

	validationByEmptyValidator, _ := noRulesValidator.Validate(password)
	validationByLengthValidator, _ := lengthValidator.Validate(password)

	suite.True(validationByEmptyValidator)
	suite.False(validationByLengthValidator)
}

func TestValidatorTestSuite(t *testing.T) {

	suite.Run(t, new(ValidatorTestSuite))
}
