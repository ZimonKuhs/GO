/*
General utility functions.

TODO:		> Should be partitioned into files according to theme.
			> Set needs its own folder. Consider new module, and hopefully generics.

@author:	Zimon Kuhs
@date:		2021-07-16
*/

package utility

type IntegerSet struct {
	set map[int]bool
}

/*** Integer set ***/

func NewIntegerSet() IntegerSet {
	return IntegerSet{
		set: map[int]bool{},
	}
}

func (set *IntegerSet) Sum() int {
	result := 0
	for key := range set.set {
		result += key
	}
	return result
}

func (set *IntegerSet) Values() []int {
	list := []int{}

	for key := range set.set {
		list = append(list, key)
	}

	return list
}

func (set *IntegerSet) Add(number int) bool {
	if _, found := set.set[number]; found {
		return false
	}

	set.set[number] = true
	return true
}

func (set *IntegerSet) Contains(number int) bool {
	return set.set[number]
}

func (set *IntegerSet) Put(number int) bool {
	if set.Contains(number) {
		return false
	}

	set.set[number] = true
	return true
}

/*** Math. ***/

/*
	Uses the sieve of Eratosthenes to find primes.

	@param ceiling	The maximum number to consider.
	@return			All primes below the input number.
*/
func Primes(ceiling int) []int {
	primes := []int{}
	result := NewIntegerSet()

	if ceiling < 2 {
		return primes
	}

	for i := 2; i < ceiling; i++ {

		if result.Contains(i) {
			continue
		}

		primes = append(primes, i)
		result.Put(i)

		for j := i; j < ceiling; j++ {
			result.Put(j)
			j *= i
		}
	}

	return primes
}

/*** String ***/

func Reverse(original string) string {
	var result string
	length := len(original)

	for i := 0; i < length; i++ {
		result += string(original[length-i-1])
	}

	return result
}
