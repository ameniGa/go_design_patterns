package visitor

import "fmt"

// Shape is the visitable object
type Shape interface {
	getType() string
	accept(Visitor)
}

// square is a concrete visitable
type square struct {
}

func (s *square) getType() string {
	return "square"
}

func (s *square) accept(visitor Visitor) {
	visitor.visitForSquare(s)
}

type circle struct {
}

func (c *circle) getType() string {
	return "circle"
}

func (c *circle) accept(visitor Visitor) {
	visitor.visitForCircle(c)
}

type triangle struct {
}

func (t *triangle) getType() string {
	return "triangle"
}

func (t *triangle) accept(visitor Visitor) {
	visitor.visitForTriangle(t)
}

type Visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForTriangle(*triangle)
}

// AreaCalculator is a concrete visitor
type AreaCalculator struct {
}

func (ac *AreaCalculator) visitForSquare(*square) {
	fmt.Println("calculating are for a square")
}

func (ac *AreaCalculator) visitForCircle(circle2 *circle) {
	fmt.Println("calculating are for a circle")
}

func (ac *AreaCalculator) visitForTriangle(*triangle) {
	fmt.Println("calculating are for a triangle")
}
