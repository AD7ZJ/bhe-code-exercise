package sieve

import (
	//"fmt"
)

type sieveImpl struct{
	sieve []bool 
}

type Sieve interface {
	NthPrime(n int64) int64
}

func NewSieve() Sieve {
    return &sieveImpl{
        sieve: make([]bool, 200),
    }
}

func (s *sieveImpl) NthPrime(n int64) int64 {
	// create an array to hold the necessary number of primes
	//s.sieve = make([]bool, n+1)
	var primes []int64        // List to store prime numbers
    var num int64 = 2         // Start with the first prime number

    // Loop until we find the n'th prime number
    for int64(len(primes)) < (n + 1) {
        // If num is marked as prime, add it to the list of primes
        if !s.sieve[num] {
            primes = append(primes, num)

            // Mark multiples of num as non-prime
            for i := num * num; i < int64(len(s.sieve)); i += num {
                s.sieve[i] = true
            }
        }

        // Move to the next number
        num++
    }

    // Return the n'th prime number
    return primes[n]




}

