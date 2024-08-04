package main

import (
	"embed"
	"io"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var assets embed.FS

var (
	audioLoaded = make(chan struct{})

	beepSound rl.Sound
	beepWave  rl.Wave
)

func init() {
	go func() {
		<-audioLoaded
		file, _ := assets.Open("assets/beep.mp3")

		b, _ := io.ReadAll(file)

		stat, _ := file.Stat()
		beepWave = rl.LoadWaveFromMemory(".mp3", b, int32(stat.Size()))
		beepSound = rl.LoadSoundFromWave(beepWave)
	}()
}

func PlayBeep() {
	rl.PlaySound(beepSound)
}
