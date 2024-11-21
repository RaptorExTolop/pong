package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type paddle struct {
	speed, width, height, X, Y, dirY int32
	playerOne                        bool
}

func (p *paddle) init(speed, width, height, X, Y int32, player bool) {
	p.speed = speed
	p.width = width
	p.height = height
	p.X = X
	p.Y = Y
	p.dirY = 0
	p.playerOne = player
}

func (p *paddle) input() {
	if p.playerOne {
		if rl.IsKeyDown(rl.KeyW) {
			p.dirY -= 1
		}
		if rl.IsKeyDown(rl.KeyS) {
			p.dirY += 1
		}
	} else {
		if rl.IsKeyDown(rl.KeyUp) {
			p.dirY -= 1
		}
		if rl.IsKeyDown(rl.KeyDown) {
			p.dirY += 1
		}
	}
}

func (p *paddle) draw() {
	rl.DrawRectangle(p.X, p.Y, p.width, p.height, rl.RayWhite)
}

func (p *paddle) update() {
	if p.dirY != 0 {
		p.Y += p.speed * p.dirY
	}
	p.dirY = 0
}
