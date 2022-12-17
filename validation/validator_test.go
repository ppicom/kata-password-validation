package validation

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidatorTestSuite struct {
	suite.Suite
}

func (suite *ValidatorTestSuite) TestValidatorShouldEnforceARuleOnAPassword() {
	password := "pass"
	noRulesValidator := NewValidator("empty_validator")
	lengthValidator := NewValidator("length_validator").With([]Rule{
		func(password string) bool { return len(password) > 5 },
	})

	validationByEmptyValidator := noRulesValidator.Validate(password)
	validationByLengthValidator := lengthValidator.Validate(password)

	suite.True(validationByEmptyValidator)
	suite.False(validationByLengthValidator)
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}
