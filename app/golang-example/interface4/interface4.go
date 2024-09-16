package main

import "fmt"

// Define the Notifier interface
type Notifier interface {
    SendNotification(message string) error
}
// EmailNotifier is a concrete implementation of Notifier
type EmailNotifier struct {
    EmailAddress string
}

func (e EmailNotifier) SendNotification(message string) error {
    fmt.Printf("Sending email to %s: %s\n", e.EmailAddress, message)
    // Simulate sending email
    return nil
}
func notify(n Notifier, message string) error {
    return n.SendNotification(message)
}
//------------------------------------------------
// SMSNotifier is another concrete implementation of Notifier
type SMSNotifier struct {
    PhoneNumber string
}

func (s SMSNotifier) SendNotification(message string) error {
    fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
    // Simulate sending SMS
    return nil
}

//-------------------------------------------------
// SlackNotifier is another concrete implementation of Notifier
type SlackNotifier struct {
    SlackChannel string
}

func (s SlackNotifier) SendNotification(message string) error {
    fmt.Printf("Sending Slack message to %s: %s\n", s.SlackChannel, message)
    // Simulate sending Slack message
    return nil
}

//--------------------------------------------

func main() {
    emailNotifier := EmailNotifier{EmailAddress: "user@example.com"}
    err := notify(emailNotifier, "Your order has been shipped!")
    if err != nil {
        fmt.Println("Failed to send notification:", err)
    }
	//--------------------------
   // emailNotifier := EmailNotifier{EmailAddress: "user@example.com"}
    smsNotifier := SMSNotifier{PhoneNumber: "+1234567890"}
    slackNotifier := SlackNotifier{SlackChannel: "#general"}

    //notify(emailNotifier, "Your order has been shipped!")
    notify(smsNotifier, "Your OTP code is 123456")
    notify(slackNotifier, "Deployment completed successfully!")

	//---------------------------
}
