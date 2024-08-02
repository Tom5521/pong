package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	rl.Vector2
	SpeedX, SpeedY float32
	Radius         float32
}

func (b *Ball) Update() {
	b.X += b.SpeedX
	b.Y += b.SpeedY

	w, h := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())

	if b.Y+b.Radius >= h || b.Y-b.Radius <= 0 {
		b.SpeedY *= -1
	}

	if b.X+b.Radius >= w || b.X-b.Radius <= 0 {
		b.SpeedX *= -1
	}
}

func (b Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), b.Radius, rl.White)
}
