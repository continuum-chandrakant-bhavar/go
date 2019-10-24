package main

import "fmt"

type person struct {
	name string
	age  int
}

func testPerson(str string) *person {

	p := person{name: str}
	p.age = 28
	return &p
}

func main() {

	fmt.Println(person{name: "Sunit", age: 28})
	fmt.Println(person{name: "Kale"})
	fmt.Println(person{"Chandu", 51})
	fmt.Println(&person{name: "Mahesh", age: 88})
	sp := person{name: "Guru", age: 66}
	spp := &sp
	fmt.Println(spp.age)
	spp.age = 89
	fmt.Println(spp.age)
	fmt.Println(sp.age)

}
