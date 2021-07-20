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
	"math"
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

/*
	Uses the sieve of erathe-atchoo.
	TODO: Move to "MATH"
*/
func compositePrimes(ceiling int) []int {
	primes := []int{}
	result := map[int]bool{}

	if ceiling < 2 {
		return primes
	}

	for i := 2; i < int(math.Sqrt(float64(ceiling))); i++ {

		// If map has key the composite number has already been marked.
		if _, ok := result[i]; ok {
			continue
		}

		primes = append(primes, i)
		result[i] = true

		for j := i; j < ceiling; j++ {
			result[j] = false
			j *= i
		}
	}

	return primes
}

/*
int np, prime[N];
bool isp[N];
void sieve(int N) {
    memset(isp, true, sizeof isp);
    isp[0] = isp[1] = false;
    for(int i=2; i<N; i++) if(isp[i]) {
        prime[++np]=i;
        for(int j=2*i; j<N; j+=i) {
            isp[j]=false;
        }
    }
}
*/

func solve(ceiling int64) (string, error) {
	factors := []int64{}
	numbers := []int64{}
	product := int64(1)

	return fmt.Sprint(sum), nil
}
