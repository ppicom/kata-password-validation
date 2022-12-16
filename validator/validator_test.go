package validator_test

import (
	"github.com/ppicom/gdc2022-go/validator"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidatorTestSuite struct {
	suite.Suite
	service validator.Service
}

func (suite *ValidatorTestSuite) Test_Acceptance() {
	service := validator.New()

	isValid := service.IsValidPassword("Ab3_56789")

	suite.True(isValid)
}

func (suite *ValidatorTestSuite) Test_Unit() {
	tests := []struct {
		name     string
		password string
		assert   func(isValid bool)
	}{
		{
			name: "given a password eight characters long, " +
				"when service::IsValidPassword is called, " +
				"then returns false",
			password: "aaaaaaaaa",
			assert: func(isValid bool) {
				suite.False(isValid)
			},
		},
		{
			name: "given a password eight characters long " +
				"and given it has a capital letter, " +
				"when service::IsValidPassword is called, " +
				"then returns false",
			password: "123A456789",
			assert: func(isValid bool) {
				suite.False(isValid)
			},
		},
		{
			name: "given a password eight characters long " +
				"and given it has a capital letter " +
				"and given it has a lowercase letter, " +
				"when service::IsValidPassword is called, " +
				"then returns false",
			password: "aaaaaaaBA",
			assert: func(isValid bool) {
				suite.False(isValid)
			},
		},
		{
			name: "given a password eight characters long " +
				"and given it has a capital letter " +
				"and given it has a lowercase letter " +
				"and given it contains a number, " +
				"when service::IsValidPassword is called, " +
				"then returns false",
			password: "123A4567b89",
			assert: func(isValid bool) {
				suite.False(isValid)
			},
		},
		{
			name: "given a password eight characters long " +
				"and given it has a capital letter " +
				"and given it has a lowercase letter " +
				"and given it contains a number " +
				"and given it contains an underscore, " +
				"when service::IsValidPassword is called, " +
				"then returns true",
			password: "123A4_567b89",
			assert: func(isValid bool) {
				suite.True(isValid)
			},
		},
	}

	for _, tt := range tests {
		suite.SetupTest()

		isValid := suite.service.IsValidPassword(tt.password)

		tt.assert(isValid)

	}
}

func (suite *ValidatorTestSuite) SetupTest() {
	suite.service = validator.New()
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}
