package main

import "fmt"

// Target interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}

// Adaptee (new payment system)
type StripePayment struct{}

func (s *StripePayment) ChargeCard(amount float64, currency string) bool {
	fmt.Printf("Charging %.2f %s via Stripe\n", amount, currency)
	return true
}

// Adapter
type StripeAdapter struct {
	stripePayment *StripePayment
}

func (sa *StripeAdapter) ProcessPayment(amount float64) bool {
	return sa.stripePayment.ChargeCard(amount, "USD")
}

// =======================================
func main() {
	adaptee := &StripePayment{}
	adapter := &StripeAdapter{stripePayment: adaptee}

	result := adapter.ProcessPayment(54)
	fmt.Println(result)

}

//=======================================
