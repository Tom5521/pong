package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Update() {
	windowWidth := float32(rl.GetScreenWidth())
	windowHeight := float32(rl.GetScreenHeight())
	if g.lastWindowSize.X != windowWidth || g.lastWindowSize.Y != windowHeight {
		g.refreshWindowSize(windowWidth, windowHeight)
	}

	if g.isWaiting4Play {
		if rl.IsKeyPressed(rl.KeySpace) {
			g.isWaiting4Play = !g.isWaiting4Play
		}
		return
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		g.isPaused = !g.isPaused
	}
	if g.isPaused {
		return
	}

	g.Player.Update()
	g.Ball.Update()
	g.CPU.Update(g.Ball.Y)

	g.checkCollisions()
	g.checkForPoints()
}

func (g *Game) refreshWindowSize(windowWidth, windowHeight float32) {
	// Update screen sizes
	g.ScreenHeight = int32(windowHeight)
	g.ScreenWidth = int32(windowWidth)

	// Precompute values.

	averageWindowSize := (windowWidth + windowHeight) / 2

	paddleHeight := (PaddleHeightPercentage / 100) * windowWidth
	paddleWidth := (PaddleWidthPercentage / 100) * windowHeight
	speed := (SpeedPercentage / 100) * averageWindowSize
	fontSize := (FontSizePercentage / 100) * averageWindowSize

	// Player
	g.Player.Width = paddleWidth
	g.Player.Height = paddleHeight
	g.Player.Speed = speed
	g.Player.X = (windowWidth - g.Player.Width) - 10

	// CPU
	g.CPU.Width = paddleWidth
	g.CPU.Height = paddleHeight
	g.CPU.Speed = speed

	// Ball
	if g.Ball.SpeedY <= -1 {
		g.Ball.SpeedY = -speed
	}
	if g.Ball.SpeedX <= -1 {
		g.Ball.SpeedX = -speed
	}
	g.Ball.Radius = (BallRadiusPercentage / 100) * averageWindowSize
	// Check if ball gets out of the window.
	if g.Ball.Y > windowHeight {
		g.Ball.Y = windowHeight - g.Ball.Radius
	}
	if g.Ball.X > windowWidth {
		g.Ball.X = windowWidth - g.Ball.Radius
	}

	// Update font sizes.
	g.pausedText.FontSize = fontSize
	g.playerPoints.FontSize = fontSize / 1.1

	g.cpuPoints.FontSize = fontSize / 1.1
	g.playText.FontSize = fontSize / 1.1

	// Update text positions.
	g.pausedText.X = (windowWidth / 2) + (100 - float32(len(g.pausedText.Text))*38)
	g.pausedText.Y = windowHeight / 3

	g.cpuPoints.Y = windowHeight / 2
	g.cpuPoints.X = (windowWidth / 2) - 100

	g.playerPoints.X = (windowWidth / 2) + (100 - 38*float32(len(g.playerPoints.Text)))
	g.playerPoints.Y = windowHeight / 2

	g.playText.X = (windowWidth / 2) - (38 * float32(len(g.playText.Text)) / 2)
	g.playText.Y = windowHeight / 4
}

func (g *Game) ResetToDefaultState() {
	g.isWaiting4Play = true

	g.Ball.X = float32(g.ScreenWidth) / 2
	g.Ball.Y = float32(g.ScreenHeight) / 2

	g.Player.Y = (float32(g.ScreenHeight) / 2) - (g.Player.Height / 2)
	g.CPU.Y = (float32(g.ScreenHeight) / 2) - (g.CPU.Height / 2)
}

func (g *Game) Pause() {
	g.isPaused = true
}

func (g *Game) UnPause() {
	g.isPaused = false
}
