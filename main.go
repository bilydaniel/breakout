package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	paddleW  = 32
	paddleH  = 8
	ballH    = 6
	ballW    = 6
	screenW  = 320
	screenH  = 240
	ballDMax = 2
	brickW   = 32
	brickH   = 8
)

type Game struct {
	Ball   *Ball
	Player *Paddle
	Bricks *[]*Brick
}

func (g *Game) Update() error {
	var dx int
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx += 1
	}

	if g.Player != nil {
		g.Player.Update(dx)
	}

	g.Ball.Update(g.Player, g.Bricks)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, "Score: "+strconv.Itoa(g.Player.Score), basicfont.Face7x13, 5, 10, color.White)
	text.Draw(screen, "Life: "+strconv.Itoa(g.Player.Lifes), basicfont.Face7x13, 5, 30, color.White)
	//TODO high score

	if g.Ball != nil {
		g.Ball.Draw(screen)
	}

	if g.Player != nil {
		g.Player.Draw(screen)

	}

	//BRICKS
	for _, brick := range *g.Bricks {
		brick.Draw(screen)
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(screenW), int(screenH)
}

func main() {
	ebiten.SetWindowSize(int(screenW*4), int(screenH*4))
	ebiten.SetWindowTitle("BREAKOUT")

	if err := ebiten.RunGame(&Game{
		Ball:   initBall(),
		Player: initPaddle(),
		Bricks: generateBricks(),
	}); err != nil {
		log.Fatal(err)
	}
}
