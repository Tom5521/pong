package main

import (
	"pong/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Update() {
	windowWidth := f(rl.GetScreenWidth())
	windowHeight := f(rl.GetScreenHeight())
	if g.lastWindowSize.X != windowWidth || g.lastWindowSize.Y != windowHeight {
		g.refreshWindowSize(windowWidth, windowHeight)
	}

	switch rl.GetKeyPressed() {
	case rl.KeyR:
		g.isWaiting4Play = true
		g.ResetToDefaultState()
		g.Player.Points = 0
		g.CPU.Points = 0

		g.playerPoints.Text = "0"
		g.cpuPoints.Text = "0"

		audio.Play(audio.Pause)
		return
	case rl.KeyM:
		audio.Mute = !audio.Mute
	case rl.KeySpace:
		if g.isWaiting4Play {
			g.isWaiting4Play = !g.isWaiting4Play
		} else {
			g.isPaused = !g.isPaused
		}
		audio.Play(audio.Pause)
	}

	if g.isPaused || g.isWaiting4Play {
		return
	}

	g.Player.Update()
	g.Ball.Update()
	g.CPU.Update(g.Ball.Y)

	g.checkCollisions()
	g.checkForPoints()
}

func (g *Game) refreshWindowSize(windowWidth, windowHeight float) {
	// Update screen sizes
	g.ScreenHeight = rint(windowHeight)
	g.ScreenWidth = rint(windowWidth)

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
	g.CPU.Speed = speed / 1.3

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

	if g.isWaiting4Play {
		g.Ball.X = windowWidth / 2
		g.Ball.Y = windowHeight / 2
	}

	// Update font sizes.
	g.pausedText.FontSize = fontSize
	g.playerPoints.FontSize = fontSize / 1.1

	g.cpuPoints.FontSize = fontSize / 1.1
	g.playText.FontSize = fontSize / 1.1
	g.mutedText.FontSize = fontSize / 1.1

	// Update text positions.
	g.pausedText.X = (windowWidth / 2) + (100 - MeasureText(g.pausedText.Text, g.pausedText.FontSize).X)
	g.pausedText.Y = windowHeight / 3

	g.cpuPoints.Y = windowHeight / 2
	g.cpuPoints.X = (windowWidth / 2) - 100

	g.playerPoints.X = (windowWidth / 2) + (100 - MeasureText(g.playerPoints.Text, g.playerPoints.FontSize).X)
	g.playerPoints.Y = windowHeight / 2

	g.playText.X = (windowWidth / 2) - (MeasureText(g.playText.Text, g.playText.FontSize).X / 2)
	g.playText.Y = windowHeight / 4
}

func (g *Game) ResetToDefaultState() {
	g.isWaiting4Play = true

	g.Ball.X = f(g.ScreenWidth) / 2
	g.Ball.Y = f(g.ScreenHeight) / 2

	g.Player.Y = (f(g.ScreenHeight) / 2) - (g.Player.Height / 2)
	g.CPU.Y = (f(g.ScreenHeight) / 2) - (g.CPU.Height / 2)
}

func (g *Game) Pause() {
	g.isPaused = true
}

func (g *Game) UnPause() {
	g.isPaused = false
}
