package main

import (
	"time"
	"strconv"
	"math/rand"
	"strings"
	"fmt"
)

func main() {
	amount := 100000000
	max := amount * 10
	min := 0
	var b strings.Builder
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < amount - 1; i++ {
		b.WriteString(strconv.Itoa(rand.Int()%(max-min) + min))
		b.WriteString(", ")
	}
	b.WriteString(strconv.Itoa(rand.Int()%(max-min) + min))
	fmt.Print(b.String())
}

