package main

type MessageSender interface {
	Send(message string) error
}

type EmailSender struct {
	// Email configuration
	name string
}

func (e *EmailSender) Send(message string) error {
	// Implementation to send an email
	return nil
}

type NotificationService struct {
	sender MessageSender
}

func NewNotificationService(sender MessageSender) *NotificationService {
	return &NotificationService{sender: sender}
}

func (ns *NotificationService) Notify(message string) error {
	return ns.sender.Send(message)
}

func main() {
	//myemail :=

}
