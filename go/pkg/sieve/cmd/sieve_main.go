package main

import (
	"fmt"
	"ssse-exercise-sieve/pkg/sieve"
)

func main() {
	sieve := sieve.NewSieve()
	n := 4

 	fmt.Printf("%dth prime number is %d\n", n, sieve.NthPrime(int64(n)))
}
