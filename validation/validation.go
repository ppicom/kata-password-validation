package validation

import "regexp"

type Service interface {
	IsValidPassword(password string, verification string) bool
}

type service struct {
	factory factory
}

func New(f factory) Service {
	return &service{
		factory: f,
	}
}

var (
	atLeastOneCapitalLetter   = regexp.MustCompile("[A-Z]+")
	atLeastOneLowercaseLetter = regexp.MustCompile("[a-z]+")
	atLeastOneNumber          = regexp.MustCompile("[1-9]+")
	atLeastOneUnderscore      = regexp.MustCompile("_+")
)

func (s service) IsValidPassword(password string, verification string) bool {
	return s.factory.Select(verification).Validate(password)
}
