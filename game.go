package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

	g.Player.Update()
	g.Ball.Update()
	g.CPU.Update(g.Ball.Y)

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
	g.Player.Width = paddleWidth
	g.Player.Height = paddleHeight
	g.Player.Speed = speed
	g.Player.X = windowWidth - g.Player.Width

	// CPU
	g.CPU.Width = paddleWidth
	g.CPU.Height = paddleHeight
	g.CPU.Speed = speed

	if g.Ball.SpeedY <= -1 {
		g.Ball.SpeedY = speed * -1
	}
	if g.Ball.SpeedX <= -1 {
		g.Ball.SpeedX = speed * -1
	}
	g.Ball.Radius = (BallRadiusPercentage / 100) * averageWindowSize
}

func (g *Game) checkCollisions() {
	// Checking for collisions
	if rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.CPU.Rectangle) {
		g.Ball.SpeedX *= -1
	}
	if rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.Player.Rectangle) {
		g.Ball.SpeedX *= -1
	}
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.Black)
	rl.DrawLine(
		g.ScreenWidth/2,
		0,
		g.ScreenWidth/2,
		g.ScreenHeight,
		rl.White,
	)
	g.Ball.Draw()
	g.Player.Draw()
	g.CPU.Draw()
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
