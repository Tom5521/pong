package audio

import (
	"embed"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var assets embed.FS

var Mute bool

var (
	Beep    rl.Sound
	Lose    rl.Sound
	Pause   rl.Sound
	Victory rl.Sound
)

func load(name string) rl.Sound {
	file, _ := assets.ReadFile("assets/" + name)

	parts := strings.Split(name, ".")

	wave := rl.LoadWaveFromMemory("."+parts[len(parts)-1], file, int32(len(file)))
	return rl.LoadSoundFromWave(wave)
}

func Play(s rl.Sound) {
	if Mute {
		return
	}
	rl.PlaySound(s)
}

func Load() {
	Beep = load("beep.mp3")
	Lose = load("lose.mp3")
	Pause = load("pause.wav")
	Victory = load("victory.mp3")
}
