package main

import "Patterns/factory/shape"

func main() {
	shape.ShapeFactory(1, 10).Area()
	shape.ShapeFactory(2, 10).Area()
}
