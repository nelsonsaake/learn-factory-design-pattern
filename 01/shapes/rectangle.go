package shapes

type rectangle struct{}

func (r rectangle) Draw() {
	draw(r)
}

func makeRectangle() rectangle {
	return rectangle{}
}
