package singleton

import (
	"sync"
)

type singleton map[string]string

type Rectangle struct {
	Len float64
}

func (r Rectangle) Area() float64 {
	return (r.Len * r.Len)
}

var (
	once     sync.Once
	onceMore sync.Once
	instance singleton
	rect     Rectangle
)

func GetInstance() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

func GetRectInstance(len float64) Rectangle {
	onceMore.Do(func() {
		rect = Rectangle{len}
	})
	return rect
}
