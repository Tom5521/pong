package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	ScreenWidth  int32
	ScreenHeight int32
	WindowTitle  string
	FPS          int32

	Ball   Ball
	Player Paddle
	CPU    CPUPaddle

	pausedText Text

	cpuPoints    Text
	playerPoints Text

	lastWindowSize rl.Vector2
	isPaused       bool
	isWaiting4Play bool
}

func NewGame(
	title string,
	fps int32,
) Game {
	return Game{
		WindowTitle:  title,
		ScreenWidth:  DefaultWidth,
		ScreenHeight: DefaultHeight,
		FPS:          fps,
		lastWindowSize: rl.Vector2{
			X: DefaultWidth,
			Y: DefaultHeight,
		},

		Ball: Ball{
			Vector2: rl.Vector2{
				X: DefaultWidth / 2,
				Y: DefaultHeight / 2,
			},
			Radius: DefaultBallRadius,
			SpeedX: DefaultSpeed,
			SpeedY: DefaultSpeed,
		},
		Player: Paddle{
			Rectangle: rl.Rectangle{
				Width:  DefaultPaddleWidth,
				Height: DefaultPaddleHeight,
				X:      DefaultWidth - DefaultPaddleWidth,
				Y:      DefaultHeight/2 - DefaultPaddleHeight/2,
			},
			Speed: DefaultSpeed,
		},
		CPU: CPUPaddle{
			Paddle: Paddle{
				Rectangle: rl.Rectangle{
					Width:  DefaultPaddleWidth,
					Height: DefaultPaddleHeight,
					X:      10,
					Y:      DefaultHeight/2 - DefaultPaddleHeight/2,
				},
				Speed: DefaultSpeed / 1.3,
			},
		},
		pausedText: Text{
			Text:     "PAUSED",
			FontSize: DefaultFontSize,
			Color:    rl.White,
		},
		cpuPoints: Text{
			Text:     "0",
			FontSize: DefaultFontSize / 1.1,
			Color:    rl.LightGray,
			Vector2: rl.Vector2{
				X: (DefaultWidth / 2) - 100,
				Y: DefaultHeight / 2,
			},
		},
		playerPoints: Text{
			Text:     "0",
			FontSize: DefaultFontSize / 1.1,
			Color:    rl.LightGray,
			Vector2: rl.Vector2{
				X: (DefaultWidth / 2) + (100 - 38),
				Y: DefaultHeight / 2,
			},
		},
	}
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.Black)
	rl.DrawLine(
		g.ScreenWidth/2,
		0,
		g.ScreenWidth/2,
		g.ScreenHeight,
		rl.White,
	)

	g.Ball.Draw()
	g.Player.Draw()
	g.CPU.Draw()

	g.playerPoints.Draw()
	g.cpuPoints.Draw()

	if g.isPaused {
		g.pausedText.Draw()
	}

	if g.isWaiting4Play {
		t := Text{
			Text:     "Press [space] to play",
			Color:    rl.White,
			FontSize: DefaultFontSize / 1.1,
			Vector2: rl.Vector2{
				Y: float32(g.ScreenHeight) / 4,
			},
		}
		totalLen := len(t.Text) * 38
		t.X = (float32(g.ScreenWidth) / 2) - float32(totalLen)/2

		t.Draw()
	}
}

func (g *Game) Create() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, g.WindowTitle)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	fmt.Println(rl.GetFontDefault().BaseSize)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}
