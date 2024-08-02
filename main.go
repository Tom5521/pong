package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	ball   Ball
	player Paddle
	cpu    CPUPaddle
)

const (
	BaseWidth  = 1280
	BaseHeight = 800

	AverageBaseSize = (float32(BaseWidth) + BaseHeight) / 2

	PaddleHeightPercentage = (120 / float32(BaseWidth)) * 100
	PaddleWidthPercentage  = (25 / float32(BaseHeight)) * 100

	SpeedPercentage      = (7 / AverageBaseSize) * 100
	BallRadiusPercentage = (20 / AverageBaseSize) * 100
)

type Game struct {
	ScreenWidth  int32
	ScreenHeight int32
	WindowTitle  string
	FPS          int32

	Ball   Ball
	Player Paddle
	CPU    CPUPaddle

	lastWindowSize rl.Vector2
}

func NewGame(
	title string,
	width, height, fps int32,
) Game {
	speed := (SpeedPercentage / 100) * AverageBaseSize
	paddleHeight := (PaddleHeightPercentage / 100) * BaseWidth
	paddleWidth := (PaddleWidthPercentage / 100) * BaseHeight

	g := Game{
		WindowTitle:  title,
		ScreenWidth:  width,
		ScreenHeight: height,
		FPS:          fps,

		Ball: Ball{
			Vector2: rl.Vector2{
				X: BaseWidth / 2,
				Y: BaseHeight / 2,
			},
			Radius: (BallRadiusPercentage / 100) * AverageBaseSize,
			SpeedX: speed,
			SpeedY: speed,
		},
		Player: Paddle{
			Rectangle: rl.Rectangle{
				Width:  paddleWidth,
				Height: paddleHeight,
				X:      BaseWidth - paddleWidth,
				Y:      BaseHeight/2 - paddleHeight/2,
			},
			Speed: speed,
		},
		CPU: CPUPaddle{
			Paddle: Paddle{
				Rectangle: rl.Rectangle{
					Width:  paddleWidth,
					Height: paddleHeight,
					X:      0,
					Y:      BaseHeight/2 - paddleHeight/2,
				},
				Speed: speed,
			},
		},
	}

	return g
}

func (g *Game) Update() {
	windowWidth := float32(rl.GetScreenWidth())
	windowHeight := float32(rl.GetScreenHeight())
	if g.lastWindowSize.X != windowWidth || g.lastWindowSize.Y != windowHeight {
		g.refreshWindowSize(windowWidth, windowHeight)
	}

	player.Update()
	ball.Update()
	cpu.Update(ball.Y)

	g.checkCollisions()
}

func (g *Game) refreshWindowSize(windowWidth, windowHeight float32) {
	g.ScreenHeight = int32(windowHeight)
	g.ScreenWidth = int32(windowWidth)

	averageWindowSize := (windowWidth + windowHeight) / 2

	paddleHeight := (PaddleHeightPercentage / 100) * windowWidth
	paddleWidth := (PaddleWidthPercentage / 100) * windowHeight
	speed := (SpeedPercentage / 100) * averageWindowSize

	// Player
	player.Width = paddleWidth
	player.Height = paddleHeight
	player.Speed = speed
	player.X = windowWidth - player.Width

	// CPU
	cpu.Width = paddleWidth
	cpu.Height = paddleHeight
	cpu.Speed = speed

	if ball.SpeedY <= -1 {
		ball.SpeedY = speed * -1
	}
	if ball.SpeedX <= -1 {
		ball.SpeedX = speed * -1
	}
	ball.Radius = (BallRadiusPercentage / 100) * averageWindowSize
}

func (g *Game) checkCollisions() {
	// Checking for collisions
	if rl.CheckCollisionCircleRec(ball.Vector2, ball.Radius, cpu.Rectangle) {
		ball.SpeedX *= -1
	}
	if rl.CheckCollisionCircleRec(ball.Vector2, ball.Radius, player.Rectangle) {
		ball.SpeedX *= -1
	}
}

func (g *Game) Draw() {
	windowWidth := int32(rl.GetScreenWidth())
	windowHeight := int32(rl.GetScreenHeight())

	rl.ClearBackground(rl.Black)
	rl.DrawLine(
		windowWidth/2,
		0,
		windowWidth/2,
		windowHeight,
		rl.White,
	)
	ball.Draw()
	player.Draw()
	cpu.Draw()
}

func (g *Game) Create() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, g.WindowTitle)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}

func main() {
	g := NewGame("Pong", BaseWidth, BaseHeight, 60)
	g.Create()
}
