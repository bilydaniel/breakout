package main

type Collider struct {
	X float32
	Y float32
	W float32
	H float32

	Right() float32
	Left() float32
	Bottom() float32
	Top() float32
}
