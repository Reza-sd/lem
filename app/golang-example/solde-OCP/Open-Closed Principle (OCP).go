package main

import "fmt"

/*
O - Open-Closed Principle (OCP)
This principle states that a struct should be open for extension but closed for modification, meaning that the behavior of a struct can be extended without changing its code. This helps to keep the code flexible and adaptable to changing requirements.



Let’s say I have a task to build a payment system that will be able to process credit card payments. It should also be flexible enough to accept different types of payment methods in future.

https://hackernoon.com/go-design-patterns-an-introduction-to-solid
*/

type PaymentMethod interface { //define behaviour (Methods) first (as interface), instead of concrete definition
	Pay()
}

type Payment struct{} //define class with main behaviour

func (p Payment) Process(pm PaymentMethod) { //<--dependency injection. (top behaviour receives an interface instead of a variable)
	pm.Pay()
}

type CreditCard struct { //concrete type of main class which has the same behaviour (method)
	amount float64
}

func (cc CreditCard) Pay() { // implement behaviour of this type
	fmt.Printf("Paid %.2f using CreditCard\n", cc.amount)
}

//-----add later-------------------------------
type PayPal struct { //another type of main class which has the same behaviour (method)
	amount float64
}

func (pp PayPal) Pay() { //another implement bahaviour of this type
	fmt.Printf("Paid %.2f using PayPal\n", pp.amount)
}

//---------------------------
//============================
func main() {
	p := Payment{}          //main object of main class
	cc := CreditCard{12.23} //the object of sub-type main class
	p.Process(cc)           // call the behaviour of the main class which recieves sub-type as argument
	//---------add later-------
	pp := PayPal{94.58}
	p.Process(pp)
}

//============================
