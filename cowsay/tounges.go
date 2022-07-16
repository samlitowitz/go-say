package cowsay

import "fmt"

type Tongue string

const (
	DeadTongue   Tongue = "U "
	StonedTongue Tongue = "U "
)

type InvalidTongueError struct {
	tongue Tongue
}

func (err *InvalidTongueError) Error() string {
	return fmt.Sprintf("Tongue must be exactly two characters long, `%s` provided.", err.tongue)
}

func IsValidTongue(tongue Tongue) (bool, error) {
	if len(tongue) != 2 {
		return false, &InvalidTongueError{tongue: tongue}
	}
	return true, nil
}
