package main

import (
	"BitArray"
	"math"
)

func generateSieve(n uint) BitArray.BitArray {
	bitArray := BitArray.NewBitArray(n)

	limit := uint(math.Floor(math.Sqrt(float64(n))))

	for i := uint(2); i < limit; i++ {
		if bitArray.Get(i) {
			continue
		}

		start := i * i
		for j := start; j < n; j += i {
			bitArray.Set(j, true)
		}
	}

	return bitArray

}

func findPrimes(bitArray BitArray.BitArray) []uint {
	output := make([]uint, 0)
	for i := uint(0); i < bitArray.Length(); i++ {
		if !bitArray.Get(uint(i)) {
			output = append(output, i)
		}
	}

	return output
}

func PrimeSieve(n uint) []uint {
	return findPrimes(generateSieve(n))[2:]
}
