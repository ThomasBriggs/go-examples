package main

import (
	_ "encoding/binary"
	"fmt"
	"os/exec"
	"strings"
)

// func splitToInt(array string, sep string) []int {
// 	var intArray []int
// 	strings.Split
// }

func main() {
	out, _ := exec.Command("arrgen", "100000000").Output()
	out = out[:len(out)-1] //removes new line char

	output := strings.Split(string(out), ", ")

	fmt.Println(output)
}
