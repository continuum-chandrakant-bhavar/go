package main

import "fmt"

type draw interface {
	circle() float32
	square() float32
}

type circle1 struct {
	radius float32
}

type area struct {
	width, lenght float32
}

func (c *circle1) circle() float32 {
	return c.radius
}

func (a *area) square() float32 {
	return a.width * a.lenght
}

func (a *area) circle() float32 {
	return a.lenght * a.width
}
func (c *circle1) square() float32 {
	return c.radius
}
func main() {
	a := area{20.0, 30.0}
	c := circle1{20.0}
	fmt.Println(a.circle())
	fmt.Println(a.square())

	fmt.Println(c.square())
	fmt.Println(c.circle())
}
