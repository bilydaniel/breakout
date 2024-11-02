package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	//TODO
	Score int
	Lifes int
	Colliding
}

func initPaddle() *Paddle {
	return &Paddle{
		Colliding: Colliding{
			X: screenW / 2,
			Y: screenH - 15,
			W: paddleW,
			H: paddleH,
		},
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

	paddle.X += float32(2 * dx)
}
