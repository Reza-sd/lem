package main
import "fmt"

//=============================
// 1. Define the Product Interface (Notifier):
// Notifier is the product interface that all concrete products will implement
type Notifier interface {
    SendNotification(message string) string
}
//==============================================
//2. Concrete Products (SMSNotifier, EmailNotifier, and PushNotifier):
// SMSNotifier is a concrete product implementing the Notifier interface
type SMSNotifier struct{}

func (s *SMSNotifier) SendNotification(message string) string {
    return fmt.Sprintf("Sending SMS notification: %s", message)
}

// EmailNotifier is another concrete product implementing the Notifier interface
type EmailNotifier struct{}

func (e *EmailNotifier) SendNotification(message string) string {
    return fmt.Sprintf("Sending Email notification: %s", message)
}

// PushNotifier is yet another concrete product implementing the Notifier interface
type PushNotifier struct{}

func (p *PushNotifier) SendNotification(message string) string {
    return fmt.Sprintf("Sending Push notification: %s", message)
}
//==============================================
//3. Factory Method (NotificationFactory):
// NotificationFactory is the factory method that creates and returns a Notifier
func NotificationFactory(notificationType string) Notifier {
    if notificationType == "sms" {
        return &SMSNotifier{}
    }
    if notificationType == "email" {
        return &EmailNotifier{}
    }
    if notificationType == "push" {
        return &PushNotifier{}
    }
    return nil
}
//==============================================
func main() {
    smsNotifier := NotificationFactory("sms")
    emailNotifier := NotificationFactory("email")
    pushNotifier := NotificationFactory("push")

    fmt.Println(smsNotifier.SendNotification("Hello via SMS!"))    // Output: Sending SMS notification: Hello via SMS!
    fmt.Println(emailNotifier.SendNotification("Hello via Email!")) // Output: Sending Email notification: Hello via Email!
    fmt.Println(pushNotifier.SendNotification("Hello via Push!"))   // Output: Sending Push notification: Hello via Push!
}
//==============================================