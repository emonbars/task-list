package main

import "fmt"

func changeValue(a int) {
	fmt.Println(&a)
	a = 36
}

func main() {
	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
}
