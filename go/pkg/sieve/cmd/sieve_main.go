package main

import (
	"fmt"
	"ssse-exercise-sieve/pkg/sieve"
)

func main() {
	sieve := sieve.NewSieve()
	n := 100

 	fmt.Printf("%dth prime number is %d\n", n, sieve.NthPrime(int64(n)))
}
