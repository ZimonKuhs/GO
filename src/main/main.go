/*
@author:	Zimon Kuhs
@date:		2021-07-16
*/

package main

import (
	"fmt"
	"os"

	"euler"
)

func main() {
	result, err := euler.Solve(1, os.Args[1:]...)
	print()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
