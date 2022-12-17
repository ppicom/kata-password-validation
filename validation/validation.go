package validation

import "regexp"

type Service interface {
	IsValidPassword(s string) bool
}

type service struct{}

func New() Service {
	return &service{}
}

var (
	atLeastOneCapitalLetter   = regexp.MustCompile("[A-Z]+")
	atLeastOneLowercaseLetter = regexp.MustCompile("[a-z]+")
	atLeastOneNumber          = regexp.MustCompile("[1-9]+")
	atLeastOneUnderscore      = regexp.MustCompile("_+")
)

func (s service) IsValidPassword(pass string) bool {
	return len(pass) > 8 &&
		atLeastOneCapitalLetter.MatchString(pass) &&
		atLeastOneLowercaseLetter.MatchString(pass) &&
		atLeastOneNumber.MatchString(pass) &&
		atLeastOneUnderscore.MatchString(pass)
}
