package shapes

import "fmt"

func Make(t Type) Shape {
	switch t {
	case Circle:
		return makeCircle()
	case Rectangle:
		return makeRectangle()
	case Square:
		return makeSquare()
	default:
		panic(fmt.Errorf("shape type not supported: %v", t))
	}
}
