package main

import rl "github.com/gen2brain/raylib-go/raylib"

func MeasureText(text string, size float) Vector {
	return rl.MeasureTextEx(
		rl.GetFontDefault(),
		text,
		size,
		5,
	)
}

type Text struct {
	Vector

	Text string

	FontSize float
	Color    rl.Color
}

func NewText(text string, pos Vector, fontSize float, color rl.Color) Text {
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
		t.Vector,
		t.FontSize,
		5,
		t.Color,
	)
}
