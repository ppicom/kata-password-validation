package validation

import "regexp"

type RuleSet interface {
	ForExpression(regex string) RuleSet
	ForLength(minimum int) RuleSet
	RunAgainst(password Password) bool
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
		rule(func(password Password) bool { return regexp.MustCompile(regex).MatchString(password.Value()) }),
	)
	return r
}

func (r *ruleset) ForLength(minimum int) RuleSet {

	r.rules = append(r.rules, rule(func(password Password) bool { return len(password.Value()) > minimum }))
	return r
}

func (r *ruleset) RunAgainst(password Password) bool {

	for _, rl := range r.rules {
		if !rl.Run(password) {
			return false
		}
	}

	return true
}
