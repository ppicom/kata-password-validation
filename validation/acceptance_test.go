package validation

import (
	"regexp"
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
	rulesForFirstValidator := []Rule{
		func(password string) bool { return len(password) > 8 },
		func(password string) bool { return regexp.MustCompile("[A-Z]+").MatchString(password) },
		func(password string) bool { return regexp.MustCompile("[a-z]+").MatchString(password) },
		func(password string) bool { return regexp.MustCompile("[0-9]+").MatchString(password) },
		func(password string) bool { return regexp.MustCompile("_+").MatchString(password) },
	}
	rulesForSecondValidator := []Rule{
		func(password string) bool { return len(password) > 6 },
		func(password string) bool { return regexp.MustCompile("[A-Z]+").MatchString(password) },
		func(password string) bool { return regexp.MustCompile("[a-z]+").MatchString(password) },
		func(password string) bool { return regexp.MustCompile("[0-9]+").MatchString(password) },
	}
	rulesForThirdValidator := []Rule{
		func(password string) bool { return len(password) > 16 },
		func(password string) bool { return regexp.MustCompile("[A-Z]+").MatchString(password) },
		func(password string) bool { return regexp.MustCompile("[a-z]+").MatchString(password) },
		func(password string) bool { return regexp.MustCompile("_+").MatchString(password) },
	}
	validators := []Validator{
		NewValidator("validation_1").With(rulesForFirstValidator),
		NewValidator("validation_2").With(rulesForSecondValidator),
		NewValidator("validation_3").With(rulesForThirdValidator),
	}
	factory := Factory().With(validators)
	service := New(factory)

	isValid := service.IsValidPassword(passLen8CapLowUnderDigit, "validation_1")
	isValidByVal2 := service.IsValidPassword(passLen6CapLower, "validation_2")
	isValidByVal3 := service.IsValidPassword(passLen16LowCapUnder, "validation_3")

	suite.Truef(isValid, "Password %s should be valid for validation 1", passLen8CapLowUnderDigit)
	suite.Truef(isValidByVal2, "Password %s should be valid for validation 2", passLen6CapLower)
	suite.Truef(isValidByVal3, "Password %s should be valid for validation 3", passLen16LowCapUnder)
}

func TestValidationTestSuite(t *testing.T) {
	suite.Run(t, new(ValidationTestSuite))
}
