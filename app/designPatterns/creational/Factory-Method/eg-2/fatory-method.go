package main
import "fmt"

//=============================
// 1. Define the Product Interface (Notifier):
// Notifier is the product interface that all concrete products will implement
type Notifier interface {
    SendNotification(message string) string
}
//==============================================