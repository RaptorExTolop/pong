package main

import (
	"fmt"

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
	b.checkCollisionWall()
	b.X += b.speed * b.dirX
	b.Y += b.speed / 3 * b.dirY
}

func (b *ball) checkCollisionWall() {
	if b.X >= screenwidth-b.radius {
		b.dirX *= -1
	}
	if b.X <= 0+b.radius {
		//score = 0
		//running = false
	}
	if b.Y >= screenHeight-b.radius || b.Y <= 0+b.radius {
		b.dirY *= -1
	}
}

func (b *ball) checkCollisionPaddle(p *paddle) {
	if p.X <= b.X && p.X+p.width >= b.X+(b.radius*2) {
		//b.dirX *= -1
		fmt.Println("x collision")
	}
	if p.Y >= b.Y && p.Y+p.height <= b.Y+(b.radius*2) {
		fmt.Println("y collision")
		b.dirY *= -1
	}
	if p.X <= b.X && p.X+p.width >= b.X+(b.radius*2) && p.Y >= b.Y && p.Y+p.height <= b.Y+(b.radius*2) {
		fmt.Println("collision")
		score++
		/*if p.X <= b.X && p.X+p.width >= b.X {
			b.dirX *= -1
		}*/
	}
}
