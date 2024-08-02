package main

const (
	BaseWidth  = 1280
	BaseHeight = 800

	AverageBaseSize = (float32(BaseWidth) + BaseHeight) / 2

	PaddleHeightPercentage = (120 / float32(BaseWidth)) * 100
	PaddleWidthPercentage  = (25 / float32(BaseHeight)) * 100

	SpeedPercentage      = (7 / AverageBaseSize) * 100
	BallRadiusPercentage = (20 / AverageBaseSize) * 100
)

func main() {
	g := NewGame("Pong", BaseWidth, BaseHeight, 1)
	g.Create()
}
