/*
'Multiples of 3 and 5'

https://projecteuler.net/problem=1

TODO: Rectify problematic argument setup to handle erroneous input generically.

@author:	Zimon Kuhs
@date:		2021-07-19
@solution:	233168
*/

package euler

import (
	"fmt"
	"strconv"
)

func euler1(args ...string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("expected one single numeric argument. Got %s", args)
	}
	number, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return "", err
	}
	if number < 0 {
		return "", fmt.Errorf("expected input to be a positive integer. Got %d", number)
	}

	return solve(number)
}

func solve(ceiling int64) (string, error) {

	return "Not yet implemented.", nil
}
