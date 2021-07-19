/*
@author:	Zimon Kuhs
@date:		2021-07-16
*/

package euler

import (
	"fmt"
)

func Solve(number int) (string, error) {
	switch number {
	case 1:
		return euler1()
	default:
		return "", fmt.Errorf("problem %d either doesn't exist or isn't implemented yet,", number)
	}
}
