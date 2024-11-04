package main

import "fmt"

//================================
// 1- Cloneable interface that all prototype objects must implement
type Cloneable interface {
	Clone() Cloneable
	GetName() string
}

//================================

// Document represents a document that can be cloned
type Document struct {
	name     string
	content  string
	authors  []string
	metadata map[string]string
}

// NewDocument creates a new document instance
func NewDocument(name, content string, authors []string, metadata map[string]string) *Document {
	return &Document{
		name:     name,
		content:  content,
		authors:  authors,
		metadata: metadata,
	}
}

// Clone creates a deep copy of the document
func (d *Document) Clone() Cloneable {
	// Create deep copies of slices and maps
	authorsCopy := make([]string, len(d.authors))
	copy(authorsCopy, d.authors)

	metadataCopy := make(map[string]string)
	for k, v := range d.metadata {
		metadataCopy[k] = v
	}

	return &Document{
		name:     d.name + "_copy",
		content:  d.content,
		authors:  authorsCopy,
		metadata: metadataCopy,
	}
}

// GetName returns the document name
func (d *Document) GetName() string {
	return d.name
}

// UpdateContent modifies the document content
func (d *Document) UpdateContent(content string) {
	d.content = content
}

// AddAuthor adds an author to the document
func (d *Document) AddAuthor(author string) {
	d.authors = append(d.authors, author)
}

// SetMetadata sets a metadata value
func (d *Document) SetMetadata(key, value string) {
	d.metadata[key] = value
}

//================================
// PrototypeRegistry manages a collection of prototype objects
type PrototypeRegistry struct {
	prototypes map[string]Cloneable
}

// NewPrototypeRegistry creates a new registry instance
func NewPrototypeRegistry() *PrototypeRegistry {
	return &PrototypeRegistry{
		prototypes: make(map[string]Cloneable),
	}
}

// Register adds a prototype to the registry
func (r *PrototypeRegistry) Register(name string, prototype Cloneable) {
	r.prototypes[name] = prototype
}

// Unregister removes a prototype from the registry
func (r *PrototypeRegistry) Unregister(name string) {
	delete(r.prototypes, name)
}

// Clone creates a copy of a registered prototype
func (r *PrototypeRegistry) Clone(name string) (Cloneable, error) {
	prototype, exists := r.prototypes[name]
	if !exists {
		return nil, fmt.Errorf("prototype with name '%s' not found", name)
	}
	return prototype.Clone(), nil
}

//================================
func main() {
	myPrototypeRegistery := NewPrototypeRegistry()
	myDcoument := NewDocument("invoce", "this is the Doc content", []string{"Reza", "Jim"}, map[string]string{"Size": "A4", "CAT": "customer"})

	myPrototypeRegistery.Register("RezaTypeDoc", myDcoument)

	myClone, err := myPrototypeRegistery.Clone("RezaTypeDoc")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cloned B %+v\n", myClone)
	}

}
