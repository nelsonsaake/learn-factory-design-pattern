package shapes

type circle struct{}

func (c circle) Draw() {
	draw(c)
}

func makeCircle() circle {
	return circle{}
}
