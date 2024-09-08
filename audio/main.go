package audio

import (
	"embed"
	"io"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var assets embed.FS

var (
	Beep    rl.Sound
	Lose    rl.Sound
	Pause   rl.Sound
	Victory rl.Sound
)

func load(name string) rl.Sound {
	file, _ := assets.Open("assets/" + name)
	bytes, _ := io.ReadAll(file)

	parts := strings.Split(name, ".")

	wave := rl.LoadWaveFromMemory("."+parts[len(parts)-1], bytes, int32(len(bytes)))
	return rl.LoadSoundFromWave(wave)
}

func Load() {
	Beep = load("beep.mp3")
	Lose = load("lose.mp3")
	Pause = load("pause.wav")
	Victory = load("victory.mp3")
}
