package validation

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type FactoryTestSuite struct {
	suite.Suite
}

func (suite *FactoryTestSuite) Test_FactoryShouldReturnTheSelectedValidator() {
	validator := NewValidator("select_me")
	validators := []Validator{
		NewValidator("select_me"),
	}
	f := Factory().With(validators)

	selectedValidator := f.Select("select_me")

	suite.True(validator.Equals(selectedValidator))
}

func TestFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}
