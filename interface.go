// package main

// import (
// 	"golangtest/gadget"
// )

// type Player interface {
// 	Play(string)
// 	Stop()
// }

// func playList(device Player, song []string) {
// 	for _, song := range song {
// 		device.Play(song)
// 	}
// 	device.Stop()

// }

// func TryOut(player Player) {
// 	player.Play("Track")
// 	player.Stop()
// 	recorder := player.(gadget.TapeRecorder)
// 	recorder.Recorder()

// }

// func main() {

// 	var player Player = gadget.TapePlayer{}
// 	mixtape := []string{ // сигмент
// 		"Супер песня !",
// 		"Не супер песня",
// 		"Еще одна хорошая песня",
// 	}
// 	playList(player, mixtape)

// 	player = gadget.TapeRecorder{}
// 	playList(player, mixtape)

// 	TryOut(gadget.TapeRecorder{})

// }
