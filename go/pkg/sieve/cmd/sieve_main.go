package main

import (
    "fmt"
    "os"
    "strconv"
    "ssse-exercise-sieve/pkg/sieve"
)

func main() {
    sieve := sieve.NewSieve()

    // Check if the correct number of arguments are provided
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run sieve_main.go <n'th prime to return>")
        return
    }

    // Parse the argument as an integer
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Invalid number:", os.Args[1])
        return
    }

    fmt.Printf("%dth prime number is %d\n", n, sieve.NthPrime(int64(n)))
}
