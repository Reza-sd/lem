package main

import "fmt"

// Abstract Products
type Button interface {
    Render() string
}

type Checkbox interface {
    Render() string
}
//--------------------------------
// Concrete Products 
type LightButton struct{}

func (b *LightButton) Render() string {
    return "Rendering a light button with white background"
}

type DarkButton struct{}

func (b *DarkButton) Render() string {
    return "Rendering a dark button with black background"
}
//---------
type LightCheckbox struct{}

func (c *LightCheckbox) Render() string {
    return "Rendering a light checkbox with white background"
}

type DarkCheckbox struct{}

func (c *DarkCheckbox) Render() string {
    return "Rendering a dark checkbox with black background"
}
//--------------------------------
// Abstract Factory interface
type UIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
}
//--------------------------------
// Concrete Factories
type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton() Button {
    return &LightButton{}
}

func (f *LightThemeFactory) CreateCheckbox() Checkbox {
    return &LightCheckbox{}
}
//-------
type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton() Button {
    return &DarkButton{}
}

func (f *DarkThemeFactory) CreateCheckbox() Checkbox {
    return &DarkCheckbox{}
}
//--------------------------------
// Client code
func createUI(thisFactory UIFactory) {
    button := thisFactory.CreateButton()
    checkbox := thisFactory.CreateCheckbox()

    fmt.Println(button.Render())
    fmt.Println(checkbox.Render())
}
//--------------------------------
func main() {
    // Create UI with light theme
    lightFactory := &LightThemeFactory{}
    fmt.Println("Creating Light Theme UI:")
    createUI(lightFactory)

    fmt.Println()

    // Create UI with dark theme
    darkFactory := &DarkThemeFactory{}
    fmt.Println("Creating Dark Theme UI:")
    createUI(darkFactory)
}
//--------------------------------