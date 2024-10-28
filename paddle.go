package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	//TODO
	X     int
	Y     int
	Score int
	Lifes int
}

func initPaddle() *Paddle {
	return &Paddle{
		X:     screenW / 2,
		Y:     screenH - 15,
		Lifes: 3,
	}
}

func (paddle *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(paddle.X), float32(paddle.Y), paddleW, paddleH, color.RGBA{0x99, 0xcc, 0xff, 0xff}, false)
}

func (paddle *Paddle) Update(dx int) {
	if paddle.X < 0 && dx < 0 {
		return
	}

	if paddle.X > screenW-paddleW && dx > 0 {
		return
	}

	paddle.X += 2 * dx
}
