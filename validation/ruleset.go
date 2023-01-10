package validation

import (
	"fmt"
	"regexp"
)

type RuleSet interface {
	ForExpression(regex string) RuleSet
	ForLength(minimum int) RuleSet
	WithUnderscore() RuleSet
	WithUppercase() RuleSet
	WithLowercase() RuleSet
	RunAgainst(password Password) (invalidBecause []string)
	CanFail(conditions int) RuleSet
}

type ruleset struct {
	rules   []Rule
	canFail int
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

	if r.canFail > len(invalidBecause) {
		invalidBecause = []string{}
	} else {
		invalidBecause = invalidBecause[r.canFail:]
	}

	return
}

// WithUnderscore implements RuleSet
func (r *ruleset) WithUnderscore() RuleSet {

	return r.ForExpression("_+")
}

// WithUppercase implements RuleSet
func (r *ruleset) WithUppercase() RuleSet {

	return r.ForExpression("[A-Z]+")
}

// WithLowercase implements RuleSet
func (r *ruleset) WithLowercase() RuleSet {

	return r.ForExpression("[a-z]+")
}

// CanFail implements RuleSet
func (r *ruleset) CanFail(conditions int) RuleSet {
	r.canFail = conditions
	return r
}
