package main

import (
	"strconv"

	"pong/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) checkCollisions() {
	// Checking for collisions

	if rl.CheckCollisionCircleRec(g.Ball.Vector, g.Ball.Radius, g.CPU.Rectangle) {
		// Check if the ball hits the top or the bottom of the paddle.
		if g.Ball.Y < g.CPU.Y || g.Ball.Y > g.CPU.Y+g.CPU.Height {
			g.Ball.SpeedY = -g.Ball.SpeedX
		}

		// Check if the ball hits the left/right side.
		if g.Ball.X < g.CPU.X || g.Ball.X > g.CPU.X+g.CPU.Width {
			g.Ball.SpeedX = -g.Ball.SpeedX
		}

		rl.PlaySound(audio.Beep)
	}
	if rl.CheckCollisionCircleRec(g.Ball.Vector, g.Ball.Radius, g.Player.Rectangle) {
		// Check if the ball hits the top or the bottom of the paddle.
		if g.Ball.Y < g.Player.Y || g.Ball.Y > g.Player.Y+g.Player.Height {
			g.Ball.SpeedY = -g.Ball.SpeedX
		}
		// Check if the ball hits the left/right side.
		if g.Ball.X < g.Player.X || g.Ball.X > g.Player.X+g.Player.Width {
			g.Ball.SpeedX = -g.Ball.SpeedX
		}

		rl.PlaySound(audio.Beep)
	}
}

func (g *Game) checkForPoints() {
	player := rl.CheckCollisionCircleRec(
		g.Ball.Vector,
		g.Ball.Radius,
		Rectangle{
			Width:  1,
			Height: f(g.ScreenHeight),
			X:      f(g.ScreenWidth),
			Y:      0,
		},
	)

	cpu := rl.CheckCollisionCircleRec(
		g.Ball.Vector,
		g.Ball.Radius,
		Rectangle{
			Width:  1,
			Height: f(g.ScreenHeight),
			X:      1,
			Y:      0,
		},
	)

	if player {
		g.CPU.Points++
		g.cpuPoints.Text = strconv.Itoa(g.CPU.Points)

		rl.PlaySound(audio.Lose)
	}

	if cpu {
		g.Player.Points++
		g.playerPoints.Text = strconv.Itoa(g.Player.Points)

		rl.PlaySound(audio.Victory)
	}

	if player || cpu {
		g.ResetToDefaultState()
	}
}
