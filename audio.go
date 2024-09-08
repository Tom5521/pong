package main

import (
	"embed"
	"io"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var assets embed.FS

var (
	beepSound    rl.Sound
	loseSound    rl.Sound
	pauseSound   rl.Sound
	victorySound rl.Sound
)

func loadSound(name string) rl.Sound {
	file, _ := assets.Open("assets/" + name)
	bytes, _ := io.ReadAll(file)

	parts := strings.SplitN(name, ".", 2)

	wave := rl.LoadWaveFromMemory("."+parts[len(parts)-1], bytes, rint(len(bytes)))
	return rl.LoadSoundFromWave(wave)
}

func loadSounds() {
	beepSound = loadSound("beep.mp3")
	loseSound = loadSound("lose.mp3")
	pauseSound = loadSound("pause.wav")
	victorySound = loadSound("victory.mp3")
}
