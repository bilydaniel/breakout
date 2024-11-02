package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	dX                   float32
	dY                   float32
	dMax                 float32
	diagonalVectorLength float32
	Colliding
}

func initBall() *Ball {
	return &Ball{
		dX:                   ballDMax,
		dY:                   ballDMax,
		dMax:                 ballDMax,
		diagonalVectorLength: float32(math.Hypot(1, 1)),
		Colliding: Colliding{
			X: float32(160 + (-20 + rand.Intn(40))),
			Y: float32(rand.Intn(120)),
			W: ballW,
			H: ballH,
		},
	}
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(ball.X), float32(ball.Y), ballH, ballW, color.RGBA{0x99, 0xcc, 0xff, 0xff}, false)
}

func (ball *Ball) Update(player *Paddle, bricks *[]*Brick) {
	if ball.dX >= 2 && ball.dY >= 2 {
		ball.dX /= float32(ball.diagonalVectorLength)
		ball.dY /= float32(ball.diagonalVectorLength)
	}
	ball.X += ball.dX
	ball.Y += ball.dY

	if ball.Y >= screenH-ballH {
		player.Lifes -= 1
		ball.Reset()
		return
	}

	if ball.Y < 0 {
		ball.dY = -ball.dY
		//TODO make the paddle smaller
	}

	if ball.X < 0 || ball.X > screenW-ballW {
		ball.dX = -ball.dX
	}

	if player != nil {
		if ball.Collides(player) {
			if ball.dY > 0 {
				ball.dY = -ball.dY
			}
		}
		//TODO optimize (probably check for distance first, check collision if close)
	}

	if bricks != nil {
		for _, brick := range *bricks {
			if ball.Collides(brick) {

			}
		}
	}
}

func (ball *Ball) Collides(collider Collider) bool {
	return ball.LeftX() < collider.RightX() &&
		ball.RightX() > collider.LeftX() &&
		ball.TopY() < collider.BottomY() &&
		ball.BottomY() > collider.TopY()
}
func (ball *Ball) Reset() {
	side := rand.Intn(3)
	coeff := 1
	if side <= 1 {
		coeff = -1
	}
	*ball = Ball{
		dX:                   float32(coeff) * ballDMax,
		dY:                   ballDMax,
		dMax:                 ballDMax,
		diagonalVectorLength: float32(math.Hypot(1, 1)),
		Colliding: Colliding{
			X: float32(160 + (-20 + rand.Intn(40))),
			Y: float32(rand.Intn(120)),
			W: ballW,
			H: ballH,
		},
	}
}
