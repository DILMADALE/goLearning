package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a)

	var b [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			b[i][j] = i + j + i + j
		}
	}

	fmt.Println(b)

	for i := 0; i < 2; i++ {
		fmt.Println(b[i])
	}
}
