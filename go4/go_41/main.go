package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&i)
	}
}

func main() {
	str1, str2 := "x", "z"
	c1 := sha256.Sum256([]byte(str1))
	c2 := sha256.Sum256([]byte(str2))
	fmt.Printf("Разница количества битов: %d\n", bitDiff(c1, c2))
}

func bitDiff(hashA, hashB [32]byte) int {
	res := 0
	for i := range hashA {
		res += int(pc[hashA[i]^hashB[i]])
	}
	return res
}
