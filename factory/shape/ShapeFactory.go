package shape

import "fmt"

type ShapeType int

const (
	DiskStorage   ShapeType = 1
	MemoryStorage ShapeType = 2
)

func ShapeFactory(t ShapeType, len float64) Shape {
	switch t {
	case MemoryStorage:
		return Square{len}
	case DiskStorage:
		return Circ{len}
	default:
		return Square{len}
	}
}

type Square struct {
	len float64
}

type Circ struct {
	rad float64
}

func (s Square) Area() {
	fmt.Println("I am a Square with Area : ", s.len*2)
}

func (c Circ) Area() {
	fmt.Println("I am a Circle with Area : ", (float64)(c.rad*c.rad*3.14))
}
