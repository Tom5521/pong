package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Text struct {
	rl.Vector2

	Text string

	FontSize float32
	Color    rl.Color
}

func NewText(text string, pos rl.Vector2, fontSize float32, color rl.Color) Text {
	return Text{
		pos,
		text,
		fontSize,
		color,
	}
}

func (t Text) Draw() {
	rl.DrawTextEx(
		rl.GetFontDefault(),
		t.Text,
		t.Vector2,
		t.FontSize,
		5,
		t.Color,
	)
}
