package shapes

type square struct{}

func (s square) Draw() {
	draw(s)
}

func makeSquare() square {
	return square{}
}
