package sieve

import (
	"fmt"
	"math"
)

type sieveImpl struct{
	primes []int64
	segmentedPrimes []int64
}

type Sieve interface {
	NthPrime(n int64) int64
}

func NewSieve() Sieve {
    return &sieveImpl{
    }
}


func (s *sieveImpl) NthPrime(n int64) int64 {
	s.SegmentedSieve(n*n)

    // Return the n'th prime number
    return append(s.primes, s.segmentedPrimes...)[n]
}



func (s *sieveImpl) RegularSieve(limit int64) {
    sieve := make([]bool, limit)
    // loop through numbers up to the square root of the limit
    for i := int64(2); i*i <= limit; i++ {
        if !sieve[i] {
            // mark multiples as non-prime
            for j := i * i; j < limit; j += i {
                sieve[j] = true // true indicates this index is not prime
            }
        }
    }

    for i := 2; i < len(sieve); i++ {
    	if !sieve[i] {
        	s.primes = append(s.primes, int64(i))
        	fmt.Printf("r%d ", i)
        }
    }
}


func (s *sieveImpl) SegmentedSieve(limit int64) {
	// divide the range into segments roughly sqrt(limit)
	segment := int64(int(math.Floor(math.Sqrt(float64(limit))) + 1))

	// calculate primes for the base range using the regular sieve
	s.RegularSieve(segment)

	// calculate upper and lower bounds for this segment
	segLo := segment
	segHi := segment * 2

	// Loop over each segment
	for segLo < limit {
		if segHi >= limit {
			segHi = limit
		}

		// allocate space for our sieve (this also clears marks from the last time)
		sieve := make([]bool, segment+1)

		// loop through all the primes found by RegularSieve() to find primes in current range
		for _, p := range s.primes {
			// Find the minimum number in this range [segLo...segHi] that is divisible (not necessarily evenly) by this prime.
			// As an example, if segLo is 11 and current prime is 3, we start with 12.
			loLim := int64(math.Floor(float64(segLo) / float64(p)) * float64(p))
			if loLim < segLo {
				loLim += p
			}

			// Mark multiples of this prime within this range as non-prime
			for j := loLim; j < segHi; j += p {
				sieve[j-segLo] = true // true indicates this index is not prime
			}
		}

		// Now collect our primes
		for i := segLo; i < segHi; i++ {
			if !sieve[i-segLo] {
				s.segmentedPrimes = append(s.segmentedPrimes, int64(i))
				fmt.Printf("%d ", i)
			}
		}

		// Update segLo and segHi for next segment
		segLo += segment
		segHi += segment
	}
}
