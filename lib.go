package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

// Random 6 digits number
func sixDigits() int64 {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}
	return n.Int64()
}

// Error Handling 
func FetchError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
