package validation

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RulesetTestSuite struct {
	suite.Suite
}

func (suite *RulesetTestSuite) TestRulesetShouldRunAllRules() {

	pwd := password("my_password")
	ruleOne, ruleTwo := NewMockRule(suite.T()), NewMockRule(suite.T())
	ruleOne.On("Run", pwd).Return(true)
	ruleTwo.On("Run", pwd).Return(true)
	ruleset := ruleset{
		rules: []Rule{
			ruleOne,
			ruleTwo,
		},
	}

	ruleset.RunAgainst(pwd)

	ruleTwo.AssertCalled(suite.T(), "Run", pwd)
	ruleOne.AssertCalled(suite.T(), "Run", pwd)
}

func TestRulesetTestSuite(t *testing.T) {
	suite.Run(t, new(RulesetTestSuite))
}
