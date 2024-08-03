package main

const (
	DefaultWidth       = 1280
	DefaultHeight      = 800
	DefaultAverageSize = (float32(DefaultWidth) + DefaultHeight) / 2

	PaddleHeightPercentage = (120 / float32(DefaultWidth)) * 100
	PaddleWidthPercentage  = (25 / float32(DefaultHeight)) * 100

	SpeedPercentage      = (7 / DefaultAverageSize) * 100
	BallRadiusPercentage = (20 / DefaultAverageSize) * 100
	FontSizePercentage   = (80 / DefaultAverageSize) * 100

	DefaultBallRadius = (BallRadiusPercentage / 100) * DefaultAverageSize
	DefaultSpeed      = (SpeedPercentage / 100) * DefaultAverageSize
	DefaultFontSize   = (FontSizePercentage / 100) * DefaultAverageSize

	DefaultPaddleHeight = (PaddleHeightPercentage / 100) * DefaultWidth
	DefaultPaddleWidth  = (PaddleWidthPercentage / 100) * DefaultHeight
)

func main() {
	g := NewGame("Pong", 60)
	g.Create()
}
