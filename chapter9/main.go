package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Rectangle : this is a rectangle
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (math.Abs(r.x1 - r.x2) + math.Abs(r.y1 - r.y2))
}

// Circle : this is a circle
type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

// Person : this is a person
type Person struct {
	Name string
}

func (p *Person) talk()  {
	fmt.Println("Hi, my name is ", p.Name)
}

// Android : Example of anonymous field
type Android struct {
	Person
	Model string
}

// Dot : Example, have no area
type Dot struct {
	x, y float64
}

func (d *Dot) area() float64 {
	return 0
}

func (d *Dot) perimeter() float64 {
	return 0
}

// Multishape : Set of shapes
type Multishape struct {
	shapes []Shape
}

func (m *Multishape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *Multishape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

// Shape : Example of interface
// Interface allows to handle different types through same name (i.e. interface)
type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}


func main() {
	c := Circle{0, 0, 10}
	r := Rectangle{0, 0, 10, 10}
	d := Dot{0, 0}
	m := Multishape{[]Shape{&c, &r, &d}}
	n := new(Multishape)
	n.shapes = m.shapes
	m.shapes = append(m.shapes, n)
	fmt.Println(totalArea(&c, &r, &d, &m))
	fmt.Println(totalPerimeter(&m))
}