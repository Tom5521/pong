package main

import (
	"embed"
	"io"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var assets embed.FS

var beepSound rl.Sound

func init() {
	rl.InitAudioDevice()
	file, _ := assets.Open("assets/beep.mp3")
	bytes, _ := io.ReadAll(file)

	beepWave := rl.LoadWaveFromMemory(".mp3", bytes, rint(len(bytes)))
	beepSound = rl.LoadSoundFromWave(beepWave)
}
