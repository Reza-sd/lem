package main

import "fmt"

// ===============================
type Computer interface {
	InsertIntoLightningPort()
}

// -----------------------
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("1-Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// -----------------------
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("2-Lightning connector is plugged into mac machine.")
}

// -----------------------
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("3-USB connector is plugged into windows machine.")
}

// -----------------------
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("4-Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

// ==================================================
func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

//==================================================
