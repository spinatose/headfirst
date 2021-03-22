package main

import (
	"ch11.com/sample/devices"
)

type Player interface {
	Play(string)
	Stop()
}

func main () {
	player := devices.TapePlayer { }
	mixtape := []string { "Are You Ok, Annie?", "Just Beat It", "Thriller" }
	playList(player, mixtape)

	recorder := devices.TapeRecorder { }
	mixtape = []string { "Billy Dean", "The Man in the Mirror", "I'm Bad" }
	playList(recorder, mixtape)

	tryOut(player)
	tryOut(recorder)
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
		device.Stop() 
	}
}

func tryOut(device Player) {
	device.Play("monkey") 
	device.Stop()
	
	recorder, ok := device.(devices.TapeRecorder)
	if ok {
		recorder.Record()
	}
}