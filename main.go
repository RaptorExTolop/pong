package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ()

var (
	screenwidth  int32 = 1200
	screenHeight int32 = 800

	running   bool     = true
	bkgColour rl.Color = rl.Lime

	score int32 = 0

	// ball
	playingBall              = ball{}
	ballSpeed          int32 = 7
	ballPosX, ballPosY int32 = screenwidth / 2, screenHeight / 2
	ballWidth          int32 = 20

	//paddles
	paddleOne = paddle{}
)

func input() {
	paddleOne.input()
}

func update() {
	running = !rl.WindowShouldClose()
	playingBall.update()
	paddleOne.update()
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColour)
	playingBall.draw()
	paddleOne.draw()
	scoreString := fmt.Sprint("Score: ", score)
	rl.DrawText(scoreString, 0, 0, 36, rl.RayWhite)

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenwidth, screenHeight, "IT'S PONG TIME")
	rl.SetTargetFPS(60)

	playingBall.init(ballSpeed, ballWidth, ballPosX, ballPosY)
	paddleOne.init(20, 25, 150, 5, (screenHeight/2)-150/2, true)
}

func main() {
	defer rl.CloseWindow()

	for running {
		input()
		update()
		draw()
	}
}
