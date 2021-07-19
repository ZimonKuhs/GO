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

/*
	I have a better idea! Start filling a list with 3^x until product exceeds ceiling. Add together.
	Then use same list, walking backwards replacing 3s with 5s, adding product if under ceiling.
	If product is above, pop a 3 from the start and keep going.
*/

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
	factors := []int64{}
	numbers := []int64{}
	product := int64(1)

	for product < ceiling {
		product *= 3
		factors = append(factors, 3)
		numbers = append(numbers, product)
	}

	pointer := len(factors) - 1
	for factors[0] != 5 {
		fmt.Sprintf("%s", factors)
		println(pointer)
		println()
		if product < ceiling {
			numbers = append(numbers, product)
			factors[pointer] = 5
			pointer--

			// *= (5 / 3); leads to float logic?
			product = product * 5 / 3
		} else {
			factors = factors[1:pointer]
			product /= 3
			pointer--
		}
	}

	sum := int64(0)
	for _, result := range numbers {
		sum += result
	}

	return fmt.Sprint(sum), nil
}
