package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ball struct {
	speed, radius, X, Y, dirX, dirY int32
}

func (b *ball) init(speed, radius, X, Y int32) {
	b.speed = speed
	b.radius = radius
	b.X = X
	b.Y = Y
	b.dirX = 1
	b.dirY = 1
}

func (b *ball) draw() {
	rl.DrawCircle(playingBall.X, playingBall.Y, float32(playingBall.radius), rl.RayWhite)
}

func (b *ball) update() {
	b.checkCollisions()
	b.X += b.speed * b.dirX
	b.Y += b.speed / 3 * b.dirY
}

func (b *ball) checkCollisions() {
	if b.X >= screenwidth-b.radius || b.X <= 0+b.radius {
		b.dirX *= -1
		score++
	}
	if b.Y >= screenHeight-b.radius || b.Y <= 0+b.radius {
		b.dirY *= -1
	}
}
