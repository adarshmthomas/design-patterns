package main

import (
	"Patterns/singletonPattern/RectShapeSingeton"
	"fmt"
)

func main() {
	s := singleton.GetInstance()
	s["this"] = "that"
	s2 := singleton.GetInstance()
	fmt.Println("This is", s2["this"])

	r := singleton.GetRectInstance(10)
	fmt.Println("Area is", r.Area())
	r1 := singleton.GetRectInstance(0)
	// 0 will not be set to length
	fmt.Println("Area is", r1.Area())
}
