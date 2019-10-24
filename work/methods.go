package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) test() int {
	return r.height * r.width
}

func (r rect) testrect() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{20, 20}
	val := r.test()
	val1 := r.testrect()
	fmt.Println("Value from method receiver", val)
	fmt.Println("Value from value method", val1)

	rp := &r
	fmt.Println("Area", rp.test())
	fmt.Println("Peri", rp.testrect())
}
