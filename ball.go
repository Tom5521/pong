package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Vector

	SpeedX, SpeedY float
	Radius         float
}

func (b *Ball) Update() {
	b.X += b.SpeedX
	b.Y += b.SpeedY

	w, h := float(rl.GetScreenWidth()), float(rl.GetScreenHeight())

	if b.Y+b.Radius >= h || b.Y-b.Radius <= 0 {
		b.SpeedY = -b.SpeedY
	}

	if b.X+b.Radius >= w || b.X-b.Radius <= 0 {
		b.SpeedX = -b.SpeedX
	}
}

func (b Ball) Draw() {
	rl.DrawCircleV(b.Vector, b.Radius, rl.White)
}
