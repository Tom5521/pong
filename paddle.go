package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	Rectangle

	Points int
	Speed  float
}

func (p Paddle) Draw() {
	rl.DrawRectangleRec(
		p.Rectangle,
		rl.White,
	)
}

func (p *Paddle) Update() {
	switch {
	case rl.IsKeyDown(rl.KeyW), rl.IsKeyDown(rl.KeyUp):
		p.Y -= p.Speed
	case rl.IsKeyDown(rl.KeyS), rl.IsKeyDown(rl.KeyDown):
		p.Y += p.Speed
	}

	p.limitMovement()
}

func (p *Paddle) limitMovement() {
	if p.Y+p.Height >= f(rl.GetScreenHeight()) {
		p.Y = f(rl.GetScreenHeight()) - p.Height
	}

	if p.Y <= 0 {
		p.Y = 0
	}
}
