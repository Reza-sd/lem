package main

import (
	"encoding/json"
	"fmt"
)

//==================================
// Prototype interface
type Cloneable interface {
	Clone() Cloneable
}

//==================================
// Concrete prototype: File
type File struct {
	Name    string
	Content string
}

func (f *File) Clone() Cloneable {
	// Perform a deep copy
	newFile := &File{}
	b, _ := json.Marshal(f)
	json.Unmarshal(b, newFile)
	return newFile
}

//----------------------
// Concrete prototype: Folder
type Folder struct {
	Name  string
	Files []Cloneable
}

func (f *Folder) Clone() Cloneable {
	newFolder := &Folder{Name: f.Name}
	for _, file := range f.Files {
		newFolder.Files = append(newFolder.Files, file.Clone())
	}
	return newFolder
}

//==================================
// Prototype registry
type FileSystem struct {
	Prototypes map[string]Cloneable
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		Prototypes: make(map[string]Cloneable),
	}
}

func (fs *FileSystem) AddPrototype(name string, prototype Cloneable) {
	fs.Prototypes[name] = prototype
}

func (fs *FileSystem) CreateFromPrototype(name string) Cloneable {
	if prototype, ok := fs.Prototypes[name]; ok {
		return prototype.Clone()
	}
	return nil
}

func main() {
	fs := NewFileSystem()

	// Add prototypes
	fs.AddPrototype("textFile", &File{Name: "sample.txt", Content: "This is a sample text file."})
	fs.AddPrototype("imageFile", &File{Name: "image.jpg", Content: "Binary image data..."})
	fs.AddPrototype("folder", &Folder{Name: "New Folder", Files: []Cloneable{
		&File{Name: "readme.txt", Content: "This is a readme file."},
	}})

	// Create objects from prototypes
	textFile := fs.CreateFromPrototype("textFile").(*File)
	imageFile := fs.CreateFromPrototype("imageFile").(*File)
	folder := fs.CreateFromPrototype("folder").(*Folder)

	// Modify the cloned objects
	textFile.Name = "newText.txt"
	imageFile.Name = "newImage.png"
	folder.Name = "Documents"

	// Print the results
	fmt.Printf("Text File: %+v\n", textFile)
	fmt.Printf("Image File: %+v\n", imageFile)
	fmt.Printf("Folder: %+v\n", folder)
	fmt.Printf("Folder's first file: %+v\n", folder.Files[0])
}
