package cowsay

import "fmt"

type Eyes string

const (
	Borg       Eyes = "=="
	DeadEyes   Eyes = "xx"
	Greedy     Eyes = "$$"
	Paranoid   Eyes = "@@"
	StonedEyes Eyes = "**"
	Tired      Eyes = "--"
	Wired      Eyes = "OO"
	Young      Eyes = ".."
)

type InvalidEyesError struct {
	eyes Eyes
}

func (err *InvalidEyesError) Error() string {
	return fmt.Sprintf("Eyes must be exactly two characters long, `%s` provided.", err.eyes)
}

func IsValidEyes(eyes Eyes) (bool, error) {
	if len(eyes) != 2 {
		return false, &InvalidEyesError{eyes: eyes}
	}
	return true, nil
}
