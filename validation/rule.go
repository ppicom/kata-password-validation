package validation

type Rule func(password string) (isValid bool)
