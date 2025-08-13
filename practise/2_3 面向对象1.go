package main

import (
	"fmt"
	"math"
)

/**
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
    考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 5}

	shapes := []Shape{rect, circle}

	for _, shape := range shapes {
		fmt.Println(shape)
		fmt.Println("面积", shape.Area())
		fmt.Println("周长", shape.Perimeter())
	}
}
