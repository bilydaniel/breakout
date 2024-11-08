package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Brick struct {
	Lifes int
	Colliding
}

func generateBricks() *[]*Brick {
	bricks := []*Brick{}
	currX := 0
	currY := 50
	go func() {
		for j := 0; j < 3; j++ {
			for i := 0; i < 10; i++ {
				bricks = append(bricks, &Brick{
					Colliding: Colliding{
						X: float32(currX),
						Y: float32(currY),
						W: brickW,
						H: brickH,
					},
					Lifes: 1,
				})
				currX += brickW
			}
			currY += 10
			currX = 0
		}
	}()
	return &bricks
}

func (brick *Brick) Draw(screen *ebiten.Image) {
	if brick != nil {
		vector.DrawFilledRect(screen, float32(brick.X), float32(brick.Y), brickW, brickH, color.RGBA{0x99, 0xcc, 0xff, 0xff}, false)
	}
}
