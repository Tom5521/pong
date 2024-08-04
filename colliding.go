package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) checkCollisions() {
	// Checking for collisions
	if rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.CPU.Rectangle) {
		g.Ball.SpeedX = -g.Ball.SpeedX
		PlayBeep()
	}
	if rl.CheckCollisionCircleRec(g.Ball.Vector2, g.Ball.Radius, g.Player.Rectangle) {
		g.Ball.SpeedX = -g.Ball.SpeedX
		PlayBeep()
	}
}

func (g *Game) checkForPoints() {
	player := rl.CheckCollisionCircleRec(
		g.Ball.Vector2,
		g.Ball.Radius,
		rl.Rectangle{
			Width:  1,
			Height: float32(g.ScreenHeight),
			X:      float32(g.ScreenWidth),
			Y:      0,
		},
	)

	cpu := rl.CheckCollisionCircleRec(
		g.Ball.Vector2,
		g.Ball.Radius,
		rl.Rectangle{
			Width:  1,
			Height: float32(g.ScreenHeight),
			X:      1,
			Y:      0,
		},
	)

	if player {
		g.CPU.Points++
		g.cpuPoints.Text = strconv.Itoa(int(g.CPU.Points))
	}
	if cpu {
		g.Player.Points++
		g.playerPoints.Text = strconv.Itoa(int(g.Player.Points))
	}

	if player || cpu {
		g.ResetToDefaultState()
	}
}
