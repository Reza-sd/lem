package main

import "fmt"

// Define a PaymentGateway interface
type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

// StripePaymentGateway is a concrete implementation of PaymentGateway
type StripePaymentGateway struct{}

func (s StripePaymentGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing payment of $%.2f using Stripe...\n", amount)
	// Simulate payment processing logic
	return nil
}

func processPayment(gateway PaymentGateway, amount float64) error {
	return gateway.ProcessPayment(amount)
}

// -----------in a 2 month=later----------------------------
// PayPalPaymentGateway is another concrete implementation of PaymentGateway
type PayPalPaymentGateway struct{}

func (p PayPalPaymentGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing payment of $%.2f using PayPal...\n", amount)
	// Simulate payment processing logic
	return nil
}

// -------------------------------------------------------
func main() {
	//---------------------------------
	stripe := StripePaymentGateway{}
	err := processPayment(stripe, 150.75)
	if err != nil {
		fmt.Println("Payment failed:", err)
	} else {
		fmt.Println("Payment successful!")
	}

	//-----------------------------
	paypal := PayPalPaymentGateway{}
	err = processPayment(paypal, 200.50)
	if err != nil {
		fmt.Println("Payment failed:", err)
	} else {
		fmt.Println("Payment successful with PayPal!")
	}

	//------------------------------------
}
