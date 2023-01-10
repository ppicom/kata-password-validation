package validation

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidationTestSuite struct {
	suite.Suite

	service Service
}

func (suite *ValidationTestSuite) Test_Acceptance() {

	const passLen6CapLower = "Aa00000"
	const passLen8CapLowUnderDigit = "Ab3_56789"
	const passLen16LowCapUnder = "aaaAAAaaaAAAaaa__"

	var firstRuleset RuleSet = NewRuleset().
		ForExpression("[A-Z]+").
		ForExpression("[A-Z]+").
		ForExpression("[A-Z]+").
		ForExpression("[A-Z]+").
		ForLength(8)

	secondRuleset := NewRuleset().
		ForLength(6).
		ForExpression("[A-Z]+").
		ForExpression("[a-z]").
		ForExpression("[0-9]+")

	thirdRuleset := NewRuleset().
		ForLength(16).
		ForExpression("[A-Z]+").
		ForExpression("[a-z]").
		ForExpression("_+")

	validators := []Validator{
		NewValidator("validation_1").With(firstRuleset),
		NewValidator("validation_2").With(secondRuleset),
		NewValidator("validation_3").With(thirdRuleset),
	}
	factory := Factory().With(validators)
	service := New(factory)

	isValid := service.IsValidPassword(passLen8CapLowUnderDigit)
	isValidByVal2 := service.IsValidPassword(passLen6CapLower)
	isValidByVal3 := service.IsValidPassword(passLen16LowCapUnder)

	suite.Truef(isValid, "Password %s should be valid for validation 1", passLen8CapLowUnderDigit)
	suite.Truef(isValidByVal2, "Password %s should be valid for validation 2", passLen6CapLower)
	suite.Truef(isValidByVal3, "Password %s should be valid for validation 3", passLen16LowCapUnder)
}

func TestValidationTestSuite(t *testing.T) {

	suite.Run(t, new(ValidationTestSuite))
}
