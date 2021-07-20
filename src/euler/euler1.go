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
	"utility"
)

func euler1(args ...string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("expected one single numeric argument. Got %s", args)
	}
	number, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return "", err
	}
	if number < 0 {
		return "", fmt.Errorf("expected input to be a positive integer. Got %d", number)
	}

	return solve(int(number), 3, 5)
}

func multiples(ceiling int, numbers ...int) map[int][]int {
	results := make(map[int][]int)

	for _, number := range numbers {
		list := []int{}

		for i := 1; i*number < ceiling; i++ {
			list = append(list, i*number)
		}

		results[number] = list
	}

	return results
}

func solve(ceiling int, numbers ...int) (string, error) {

	set := utility.NewIntegerSet()
	for _, list := range multiples(ceiling, numbers...) {
		for _, value := range list {
			set.Add(value)
		}
	}

	return fmt.Sprint(set.Sum()), nil
}
