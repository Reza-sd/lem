package main

import "fmt"

//=================================================
/*
The Composite pattern is a structural design pattern that allows you to compose objects into tree structures to represent part-whole hierarchies. In Golang, this pattern is implemented using interfaces and struct composition.

Key Components
1- Component: An interface that declares common operations for both simple and complex objects.
2- Leaf: Represents end objects of a composition with no children.
3- Composite: Stores child components and implements component-related operations.

*/

//=================================================
// Component interface
type FileSystemComponent interface {
	Display(indent string)
}

//=================================================
// Leaf
type File struct {
	name string
}

func (f *File) Display(indent string) {
	fmt.Println(indent + f.name)
}

// Composite
type Directory struct {
	name     string
	children []FileSystemComponent
}

func (d *Directory) Display(indent string) {
	fmt.Println(indent + d.name)
	for _, child := range d.children {
		child.Display(indent + "  ")
	}
}

//=================================================
func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

//=================================================

//=================================================
func main() {
	root := &Directory{name: "Root"}

	documents := &Directory{name: "Documents"}
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}
	documents.Add(file1)
	documents.Add(file2)

	pictures := &Directory{name: "Pictures"}
	file3 := &File{name: "pic1.jpg"}
	pictures.Add(file3)

	root.Add(documents)
	root.Add(pictures)

	root.Display("")
}

//=================================================
