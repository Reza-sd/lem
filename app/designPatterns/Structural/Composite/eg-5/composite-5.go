package main

import "fmt"

//========================================
/*
The Composite Pattern is a structural design pattern used to treat individual objects and compositions of objects uniformly. In other words, it allows clients to work with both individual objects (leaf nodes) and groups of objects (composite nodes) in the same way, which simplifies client code that deals with tree-like structures.

Key Concepts in Composite Pattern:

Component: The interface or abstract class that defines the operations applicable to both leaf and composite objects.

Leaf: The individual object that doesn't have any children (i.e., it represents a leaf node in the tree).

Composite: A composite object that holds child components and implements the component interface.

Composite Pattern in Golang:

Since Go does not have classes or inheritance, we implement the Composite Pattern using interfaces and struct composition. Here’s how we can implement the pattern in Golang.
*/

//========================================

//1. Define the Component Interface
//The component interface defines the operations that will be common to both leaves and composite objects.

// Component interface
type Component interface {
	Play()
	GetName() string
}

//========================================
//2. Define the Leaf Objects
//These are individual objects that implement the Component interface.

// Leaf 1: Song
type Song struct {
	name string
}

func (s *Song) Play() {
	fmt.Printf("Playing song: %s\n", s.name)
}

func (s *Song) GetName() string {
	return s.name
}

// Leaf 2: Podcast
type Podcast struct {
	name string
}

func (p *Podcast) Play() {
	fmt.Printf("Playing podcast: %s\n", p.name)
}

func (p *Podcast) GetName() string {
	return p.name
}

//========================================
//3. Define the Composite Object
//The composite object contains a collection of components (both leaves and other composites) and implements the Component interface.

// Composite: Playlist
type Playlist struct {
	name       string
	components []Component
}

func (p *Playlist) Play() {
	fmt.Printf("Playing playlist: %s\n", p.name)
	for _, component := range p.components {
		component.Play()
	}
}

func (p *Playlist) Add(c Component) {
	p.components = append(p.components, c)
}

func (p *Playlist) Remove(c Component) {
	for i, component := range p.components {
		if component.GetName() == c.GetName() {
			p.components = append(p.components[:i], p.components[i+1:]...)
			break
		}
	}
}

func (p *Playlist) GetName() string {
	return p.name
}

//=========================================
/*
The Composite Pattern is useful when dealing with tree-like structures where you want to treat both individual objects and groups of objects in the same way. In Go, this can be achieved using interfaces and composition, as demonstrated with songs, podcasts, and playlists. This pattern helps simplify client code by allowing uniform treatment of both simple and composite objects.
*/

// ========================================
func main() {
	// Create leaf components
	song1 := &Song{name: "Song A"}
	song2 := &Song{name: "Song B"}
	podcast1 := &Podcast{name: "Podcast X"}

	// Create a playlist (composite)
	playlist1 := &Playlist{name: "Morning Playlist"}

	// Add leaf components to the playlist
	playlist1.Add(song1)
	playlist1.Add(podcast1)

	// Create another playlist (composite)
	playlist2 := &Playlist{name: "Evening Playlist"}
	playlist2.Add(song2)

	// Add one playlist inside another playlist (composite inside composite)
	playlist1.Add(playlist2)

	// Play the playlists
	playlist1.Play()
}

//============================================
