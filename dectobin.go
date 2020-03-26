package main

import (
	"fmt"
	"strconv"
)

func main() {
	var decimal int64
	fmt.Print("Enter Decimal Number: ")
	fmt.Scanln(&decimal)
	output := strconv.FormatInt(decimal, 2)
	fmt.Print("Output: ", output)

}
