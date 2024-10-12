package main

import "fmt"

/*
The Adapter Pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate. It works by wrapping one of the objects to translate its interface into a form that another object can use.

In Go, we can use interfaces to implement the Adapter pattern, where an adapter acts as a wrapper between two incompatible interfaces. The adapter implements the interface expected by the client, while internally converting calls to the interface of the adaptee (the object we are adapting).
//====================
Example: Adapter Pattern in Go
Let's say we have a MediaPlayer interface that only supports playing MP3 files, but we also have an advanced media player that can play MP4 and VLC files. To make the AdvancedMediaPlayer compatible with the MediaPlayer interface, we will use the Adapter pattern.
//===================
*/
//==============================================
//1. Define the Target Interface (MediaPlayer):
//The MediaPlayer interface is what our client expects to work with. It has a method PlayMP3().

// MediaPlayer is the target interface that the client expects to use
type MediaPlayer interface {
	PlayMP3(file string)
}

//==============================================
//2. Create the Adaptee (AdvancedMediaPlayer):
//The AdvancedMediaPlayer interface represents the more advanced media player that can play MP4 and VLC formats.

// AdvancedMediaPlayer is the adaptee interface with more advanced features
type AdvancedMediaPlayer interface {
	Play(file string)
	//PlayVLC(file string)
}

// MP4Player is a concrete adaptee that can play MP4 files
type MP4Player struct{}

func (p *MP4Player) Play(file string) {
	fmt.Printf("Playing MP4 file: %s\n", file)
}

// VLCPlayer is another concrete adaptee that can play VLC files
type VLCPlayer struct{}

func (p *VLCPlayer) Play(file string) {
	fmt.Printf("Playing VLC file: %s\n", file)
}

//==============================================
//3. Create the Adapter (MediaAdapter):
//The MediaAdapter will implement the MediaPlayer interface, but internally it will delegate the calls to the AdvancedMediaPlayer. This way, it "adapts" the interface.

// MediaAdapter adapts AdvancedMediaPlayer to the MediaPlayer interface
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

// NewMediaAdapter creates a new MediaAdapter for the given format
func NewMediaAdapter(format string) MediaPlayer {
	if format == "mp4" {
		return &MediaAdapter{advancedPlayer: &MP4Player{}}
	} else if format == "vlc" {
		return &MediaAdapter{advancedPlayer: &VLCPlayer{}}
	}
	return nil
}

// PlayMP3 adapts the PlayMP3 method to the advanced media player methods
func (adapter *MediaAdapter) PlayMP3(file string) {
	if mp4Player, ok := adapter.advancedPlayer.(*MP4Player); ok {
		mp4Player.Play(file)
	} else if vlcPlayer, ok := adapter.advancedPlayer.(*VLCPlayer); ok {
		vlcPlayer.Play(file)
	}
}

//==============================================

//4. Create the Client (AudioPlayer):

// AudioPlayer is the client that plays only MP3 files directly
type AudioPlayer struct{}

func (p *AudioPlayer) Play(format string, file string) {
	if format == "mp3" {
		fmt.Printf("Playing MP3 file: %s\n", file)
	} else if format == "mp4" || format == "vlc" {
		adapter := NewMediaAdapter(format)
		adapter.PlayMP3(file)

	} else {
		fmt.Println("Invalid media format:", format)
	}
}

// ==============================================
func main() {
	player := &AudioPlayer{}

	// Play MP3 file directly
	player.Play("mp3", "song.mp3")

	// Use adapter to play MP4 file
	player.Play("mp4", "movie.mp4")

	// Use adapter to play VLC file
	player.Play("vlc", "video.vlc")

	// Try playing an unsupported format
	player.Play("avi", "clip.avi")
}

//==============================================
