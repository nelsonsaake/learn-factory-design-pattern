package main

import "github.com/nelsonsaake/dps/factory/shapes"

func main() {

	// make a circle and draw it
	c := shapes.Make("circle")
	c.Draw()

	// make rectangle and draw it
	r := shapes.Make("rectangle")
	r.Draw()

	// make square and draw it
	s := shapes.Make("square")
	s.Draw()
}
