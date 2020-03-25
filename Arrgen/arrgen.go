package main

import (
	"strings"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func printRandomArray(params []string) {
	var amount, min, max int
	var err error

	switch len(params) {
	case 1:
		amount, err = strconv.Atoi(params[0])
		max = amount * 10
		min = 0
	case 2:
		amount, err = strconv.Atoi(params[0])
		max, err = strconv.Atoi(params[1])
		min = 0
	case 3:
		amount, err = strconv.Atoi(params[0])
		min, err = strconv.Atoi(params[1])
		max, err = strconv.Atoi(params[2])
	default:
		fmt.Println(err)
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	var sb strings.Builder
	for i := 0; i < amount-1; i++ {
		sb.WriteString(strconv.Itoa((rand.Int() % (max - min)) + min))
		sb.WriteString(", ")
	}
	sb.WriteString(strconv.Itoa((rand.Int() % (max - min)) + min))
	fmt.Println(sb.String())
}

func consecArray(params []string) []int {
	var amount, start int
	var e error

	switch len(params) {
	case 1:
		amount, e = strconv.Atoi(params[0])
		start = 0
	case 2:
		amount, e = strconv.Atoi(params[0])
		start, e = strconv.Atoi(params[1])
	default:
		fmt.Println(e)
	}

	var arr []int

	for i := 0; i < amount; i++ {
		arr = append(arr, i+start)
	}

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < amount*10; i++ {
		swap(arr, rand.Int()%amount, rand.Int()%amount)
	}

	return arr
}

func swap(arr []int, index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

func formatArr(arr []int) string {
	var o string
	for i := 0; i < len(arr)-1; i++ {
		o += strconv.Itoa(arr[i]) + ", "
	}
	o += strconv.Itoa(arr[len(arr)-1])
	return o
}

func main() {
	consecutive := flag.Bool("c", false, "The array will be full of consecutive in a random order")
	flag.Parse()

	if *consecutive {
		fmt.Println(formatArr(consecArray(flag.Args())))
	} else {
		printRandomArray(flag.Args())
	}
}
