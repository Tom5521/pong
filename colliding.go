package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) checkCollisions() {
	// Checking for collisions
	if rl.CheckCollisionCircleRec(g.Ball.Vector, g.Ball.Radius, g.CPU.Rectangle) {
		g.Ball.SpeedX = -g.Ball.SpeedX
		rl.PlaySound(beepSound)
	}
	if rl.CheckCollisionCircleRec(g.Ball.Vector, g.Ball.Radius, g.Player.Rectangle) {
		g.Ball.SpeedX = -g.Ball.SpeedX
		rl.PlaySound(beepSound)
	}
}

func (g *Game) checkForPoints() {
	player := rl.CheckCollisionCircleRec(
		g.Ball.Vector,
		g.Ball.Radius,
		Rectangle{
			Width:  1,
			Height: float(g.ScreenHeight),
			X:      float(g.ScreenWidth),
			Y:      0,
		},
	)

	cpu := rl.CheckCollisionCircleRec(
		g.Ball.Vector,
		g.Ball.Radius,
		Rectangle{
			Width:  1,
			Height: float(g.ScreenHeight),
			X:      1,
			Y:      0,
		},
	)

	if player {
		g.CPU.Points++
		g.cpuPoints.Text = strconv.Itoa(g.CPU.Points)
	}
	if cpu {
		g.Player.Points++
		g.playerPoints.Text = strconv.Itoa(g.Player.Points)
	}

	if player || cpu {
		g.ResetToDefaultState()
	}
}
