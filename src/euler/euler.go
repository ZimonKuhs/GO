package euler

import (
	"errors"
	"fmt"
)

func Solve(number int) {
	switch number {
	case 12:
		return euler12.Solve()
	default:
		return -1, errors.New(fmt.Sprintf("Problem %d either doesn't exist or isn't implemented yet,", number))
	}
	return -1
}
