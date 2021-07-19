/*
General utility functions.

TODO:		Should be partitioned into files according to theme.

@author:	Zimon Kuhs
@date:		2021-07-16
*/

package utility

func Reverse(original string) string {
	var result string
	length := len(original)

	for i := 0; i < length; i++ {
		result += string(original[length-i-1])
	}

	return result
}
