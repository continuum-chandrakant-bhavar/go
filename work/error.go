package main

import (
	"errors"
	"fmt"
)

func f1(a int) (int, error) {
	if a == 42 {
		return -1, errors.New("error in f1")
	}

	return a, nil
}

type argError struct {
	arg  int
	prob string
}

func (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}

func f2(a int) (int, error) {
	if a == 42 {
		return -1, &argError{a, "Not accapeted"}
	}

	return a + 3, nil
}

func main() {
	for _, i := range []int{1, 42} {
		if res, err := f1(i); err != nil {
			fmt.Println("Work Fail", err)
		} else {
			fmt.Println("Worked fine", res)
		}
	}

	for _, a := range []int{1, 42} {
		if res1, err := f2(a); err != nil {
			fmt.Println("Work failed", err)
		} else {
			fmt.Println(" Work fine", res1)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
