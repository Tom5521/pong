package main

import (
	"pong/audio"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	ScreenWidth  rint
	ScreenHeight rint
	WindowTitle  string
	FPS          rint

	Ball   Ball
	Player Paddle
	CPU    CPUPaddle

	pausedText Text
	playText   Text
	mutedText  Text

	cpuPoints    Text
	playerPoints Text

	lastWindowSize Vector
	isPaused       bool
	isWaiting4Play bool
}

func NewGame(
	title string,
	fps rint,
) Game {
	g := Game{
		WindowTitle:  title,
		ScreenWidth:  DefaultWidth,
		ScreenHeight: DefaultHeight,
		FPS:          fps,
		lastWindowSize: Vector{
			X: DefaultWidth,
			Y: DefaultHeight,
		},

		Ball: Ball{
			Vector: Vector{
				X: DefaultWidth / 2,
				Y: DefaultHeight / 2,
			},
			Radius: DefaultBallRadius,
			SpeedX: DefaultSpeed,
			SpeedY: DefaultSpeed,
		},
		Player: Paddle{
			Rectangle: Rectangle{
				Width:  DefaultPaddleWidth,
				Height: DefaultPaddleHeight,
				X:      (DefaultWidth - DefaultPaddleWidth) - 10,
				Y:      DefaultHeight/2 - DefaultPaddleHeight/2,
			},
			Speed: DefaultSpeed,
		},
		CPU: CPUPaddle{
			Paddle: Paddle{
				Rectangle: Rectangle{
					Width:  DefaultPaddleWidth,
					Height: DefaultPaddleHeight,
					X:      10,
					Y:      DefaultHeight/2 - DefaultPaddleHeight/2,
				},
				Speed: DefaultSpeed / 1.3,
			},
		},
		isWaiting4Play: true,
	}

	g.pausedText = Text{
		Text:     "PAUSED",
		FontSize: DefaultFontSize,
		Color:    rl.White,
	}
	g.pausedText.X = (DefaultWidth / 2) + (100 - MeasureText(g.pausedText.Text, DefaultFontSize).X)
	g.pausedText.Y = DefaultHeight / 2

	g.playerPoints = Text{
		Text:     "0",
		FontSize: DefaultFontSize / 1.1,
		Color:    rl.LightGray,
	}
	g.playerPoints.X = (DefaultWidth / 2) + (100 - MeasureText(g.playerPoints.Text, g.playerPoints.FontSize).X)
	g.playerPoints.Y = DefaultHeight / 2

	g.cpuPoints = Text{
		Text:     "0",
		FontSize: DefaultFontSize / 1.1,
		Color:    rl.LightGray,
	}
	g.cpuPoints.X = (DefaultWidth / 2) + (100 - MeasureText(g.cpuPoints.Text, g.cpuPoints.FontSize).X)
	g.cpuPoints.Y = DefaultHeight / 2

	g.playText = Text{
		Text:     "Press [space] to play",
		Color:    rl.White,
		FontSize: DefaultFontSize / 1.1,
		Vector: Vector{
			Y: f(g.ScreenHeight) / 4,
		},
	}
	g.playText.X = (DefaultWidth / 2) - MeasureText(g.playText.Text, g.playText.FontSize).X

	g.mutedText = Text{
		Text:     "MUTED",
		Color:    rl.White,
		FontSize: DefaultFontSize / 1.1,
	}

	return g
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
		g.playText.Draw()
	}

	if audio.Mute {
		g.mutedText.Draw()
	}
}

func (g *Game) CreateLoop() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, g.WindowTitle)
	rl.InitAudioDevice()
	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()
	rl.SetTargetFPS(60)

	audio.Load()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}
