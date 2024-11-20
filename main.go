package main

import rl "github.com/gen2brain/raylib-go/raylib"

const ()

var (
	screenwidth  int32 = 1280
	screenHeight int32 = 720

	running   bool     = true
	bkgColour rl.Color = rl.Lime
)

func input() {}

func update() {
	running = !rl.WindowShouldClose()
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColour)

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenwidth, screenHeight, "IT'S PONG TIME")
	rl.SetTargetFPS(60)
}

func main() {
	defer rl.CloseWindow()

	for running {
		input()
		update()
		draw()
	}
}
