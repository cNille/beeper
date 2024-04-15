package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func main() {
  // Open the mp3 file
  f, err := os.Open("notification.mp3")
  if err != nil {
    log.Fatal(err)
  }

	// Decode the WAV file from the embedded bytes
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	// Initialize the speaker
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Play the beep, use channel to wait for the end of the playback
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Wait for the end of the playback
	<-done
}

