package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	DefaultWidth       = 1280
	DefaultHeight      = 800
	DefaultAverageSize = (f(DefaultWidth) + DefaultHeight) / 2

	PaddleHeightPercentage = (120 / f(DefaultWidth)) * 100
	PaddleWidthPercentage  = (25 / f(DefaultHeight)) * 100

	SpeedPercentage      = (7 / DefaultAverageSize) * 100
	BallRadiusPercentage = (20 / DefaultAverageSize) * 100
	FontSizePercentage   = (80 / DefaultAverageSize) * 100

	DefaultBallRadius = (BallRadiusPercentage / 100) * DefaultAverageSize
	DefaultSpeed      = (SpeedPercentage / 100) * DefaultAverageSize
	DefaultFontSize   = (FontSizePercentage / 100) * DefaultAverageSize

	DefaultPaddleHeight = (PaddleHeightPercentage / 100) * DefaultWidth
	DefaultPaddleWidth  = (PaddleWidthPercentage / 100) * DefaultHeight
)

type (
	float = float32
	// This will be used **ONLY FOR TYPE CONVERSIONS**.
	f         = float
	rint      = int32
	Vector    = rl.Vector2
	Rectangle = rl.Rectangle
)

func main() {
	g := NewGame("Pong", 60)
	g.CreateLoop()
}
