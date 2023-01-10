package validation

import (
	"fmt"
	"regexp"
)

type RuleSet interface {
	ForExpression(regex string) RuleSet
	ForLength(minimum int) RuleSet
	RunAgainst(password Password) (invalidBecause []string)
}

type ruleset struct {
	rules []Rule
}

func NewRuleset() RuleSet {

	return &ruleset{
		rules: make([]Rule, 0),
	}
}

func (r *ruleset) ForExpression(regex string) RuleSet {

	r.rules = append(
		r.rules,
		rule(func(password Password) (invalidBecause string) {
			if !regexp.MustCompile(regex).MatchString(password.Value()) {
				switch regex {
				case "[A-Z]+":
					invalidBecause = "Password lacks an uppercase letter."
				case "[a-z]":
					invalidBecause = "Password lacks a lowercase letter."
				case "_+":
					invalidBecause = "Password lacks an underscore."
				}
			}

			return
		}),
	)
	return r
}

func (r *ruleset) ForLength(minimum int) RuleSet {

	r.rules = append(r.rules, rule(func(password Password) (invalidBecause string) {
		if len(password.Value()) <= minimum {
			invalidBecause = fmt.Sprintf("Required len is %d", minimum)
		}

		return
	}))
	return r
}

func (r *ruleset) RunAgainst(password Password) (invalidBecause []string) {

	for _, rl := range r.rules {

		reason := rl.Run(password)
		if reason != "" {

			invalidBecause = append(invalidBecause, reason)
		}
	}

	return
}
