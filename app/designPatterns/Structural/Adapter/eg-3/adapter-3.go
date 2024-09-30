package main

import (
	"fmt"
	"strings"
)

//=================================================
/*
The Adapter pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate. It's particularly useful when you want to use an existing class, but its interface isn't compatible with the rest of your code.
*/
//======================================================
// Target interface
type TextFormatter interface {
	FormatText(text string) string
}

// Adaptee (the class that needs adapting)
type LegacyTextFormatter struct{}

func (l *LegacyTextFormatter) FormatToLowerCase(text string) string {
	return strings.ToLower(text)
}

// Adapter
type LowerCaseFormatter struct {
	legacyFormatter *LegacyTextFormatter
}

func NewLowerCaseFormatter() *LowerCaseFormatter {
	return &LowerCaseFormatter{
		legacyFormatter: &LegacyTextFormatter{},
	}
}

func (a *LowerCaseFormatter) FormatText(text string) string {
	return a.legacyFormatter.FormatToLowerCase(text)
}

// Another Adaptee
type ReverseFormatter struct{}

func (r *ReverseFormatter) ReverseString(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Another Adapter
type ReverseAdapter struct {
	reverseFormatter *ReverseFormatter
}

func NewReverseAdapter() *ReverseAdapter {
	return &ReverseAdapter{
		reverseFormatter: &ReverseFormatter{},
	}
}

func (a *ReverseAdapter) FormatText(text string) string {
	return a.reverseFormatter.ReverseString(text)
}

// Client code
func ClientCode(formatter TextFormatter) {
	result := formatter.FormatText("Hello, Adapter Pattern!")
	fmt.Println(result)
}

//=================================================
func main() {
	fmt.Println("Client code with LowerCaseFormatter:")
	lowerCaseFormatter := NewLowerCaseFormatter()
	ClientCode(lowerCaseFormatter)

	fmt.Println("\nClient code with ReverseAdapter:")
	reverseAdapter := NewReverseAdapter()
	ClientCode(reverseAdapter)
}

//=================================================
