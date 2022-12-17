package validation

type Password interface {
	Value() string
}

type password string

func (p password) Value() string {
	return string(p)
}
