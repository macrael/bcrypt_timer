package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("HELLO TIMER")

	// this function attempts to time a certain number of iterations of bcrypt.

	testSSN := "123-45-6789"
	testBytes := []byte(testSSN)

	testWorkFactors := []int{8, 10, 12, 14, 16}

	for _, testWorkFactor := range testWorkFactors {

		attemptCount := 100
		start := time.Now()
		for i := 0; i < attemptCount; i++ {
			_, err := bcrypt.GenerateFromPassword(testBytes, testWorkFactor) // -1 chooses the default cost
			if err != nil {
				fmt.Println("Got an error!")
				break
			}
		}
		elapsed := time.Since(start)

		hps := float64(attemptCount) / (float64(elapsed) / float64(time.Second))

		fmt.Printf("Took %v elapsed to generate %v hashes with work factor %v (%v hash/second)\n", elapsed, attemptCount, testWorkFactor, hps)
	}
}
