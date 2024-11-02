package main

type Collider interface {
	RightX() float32
	LeftX() float32
	BottomY() float32
	TopY() float32
}
type Colliding struct {
	X float32
	Y float32
	W float32
	H float32
}

func (colliding *Colliding) RightX() float32 {
	return colliding.X + colliding.W
}

func (colliding *Colliding) LeftX() float32 {
	return colliding.X
}

func (colliding *Colliding) BottomY() float32 {
	return colliding.Y + colliding.H
}

func (colliding *Colliding) TopY() float32 {
	return colliding.Y
}
