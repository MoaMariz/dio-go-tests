package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 3)
	z := diminui(100, 20)
	w := divide(100, 25)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}

	return total

}

func diminui(a int, i ...int) int {
	total := a
	for _, v := range i {
		total -= v
	}

	return total

}

func divide(a int, i ...int) int {
	total := a
	for _, v := range i {
		total = total / v
	}

	return total

}
