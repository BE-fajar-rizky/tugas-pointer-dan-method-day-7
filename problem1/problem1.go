package main

import "fmt"

func Swap(a, b *int) (int, int) {

	var nilaia int = *a
	var nilaib int = *b
	var swap int = nilaia

	nilaia = nilaib
	nilaib = nilaib - swap

	return nilaia, nilaib

}

func main() {
	a := 10
	b := 20

	Swap(&a, &b)
	fmt.Println(a, b)
}
