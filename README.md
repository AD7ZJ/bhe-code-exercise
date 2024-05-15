# My implementation

I chose to implement this in go to try and familiarize myself with the language (my background is embedded c programming). After getting the regular sieve working, I then implemented a segmented sieve to reduce memory usage for large primes. And lastly, I figured out a way to estimate the range needed to calculate primes over in order to come up with the n'th prime. 

Files of interest:
- [sieve.go](go/pkg/sieve/sieve.go)
- [sieve_test.go](go/pkg/sieve/sieve_test.go)
- [sieve_main.go](go/pkg/sieve/cmd/sieve_main.go)

You can run the implementation from command line inside the go package, using the following (I used this a lot for figuring out the implementation): 

`go run cmd/sieve_main.go <n>`

Or alternatively, kick off the tests:

`go test`

---
<sup>Original README.md<sup/>
# BHE Software Engineer Coding Exercise

## The Sieve of Eratosthenes

Prime numbers have many modern day applications and a long history in mathematics. Utilizing your own resources research the sieve of Eratosthenes, an algorithm for generating prime numbers. Based on your research, implement an API that allows the caller to retrieve the Nth prime number.
Some stub code and a test suite have been provided as a convenience, however, you are encouraged to deviate from Eratosthenes's algorithm, modify the existing functions/methods or anything else that might showcase your ability provided the following requirements are satisfied.
Stub code has been provided in Go, C#, and Javascript. Please use the language that is most appropriate based on your own skillset

### Requirements

- Fork this repo to implement your solution
- The library package provides an API for retrieving the Nth prime number using 0-based indexing where the 0th prime number is 2
- Interviewers must be able to execute a suite of tests
  - Go: `go test ./...`
  - C#: `dotnet test Sieve.Tests`
  - Javascript: `npm run test`
- Your solution is committed to your project's `main` branch, no uncommitted changes or untracked files please
- Submit the link to your public fork for review

### Considerations

During the technical interview, your submission will be discussed, and you will be evaluated in the following areas:

- Technical ability
- Communication skills
- Work habits and complementary skills
