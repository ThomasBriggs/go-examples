package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	parsedInt, err := strconv.ParseUint(args[0], 10, 0)
	if err != nil {
		panic("Please enter a integer above 0!")
	}
	primes := PrimeSieve(uint(parsedInt))
	fmt.Println(len(primes))

}
