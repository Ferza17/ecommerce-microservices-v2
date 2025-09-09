package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInvoiceCode() string {
	// Seed the random number generator to ensure different results each time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number in the range [1000000000, 1999999999] (10 digits)
	randomNumber := rand.Intn(1000000000) + 1000000000

	// Format the invoice code
	invoiceCode := fmt.Sprintf("INV/2025/%d", randomNumber)

	return invoiceCode
}
