package main

import (
	"fmt"
	"os"
	"strconv"
)

// const m float64 = 1.422349587101

func main() {
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum : ", sum)
	fname := "Original"
	fmt.Println("Name : ", fname)
	updateName(&fname)
	fmt.Println("Name : ", fname)
}

func sum(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return result
}

func updateName(name *string) {
	*name = "Changed"
}

// func main() {
// 	x, y, z := 10, 15.5, "Gophers"
// 	score := []int{10, 20, 30}

// 	fmt.Printf("X : %d, Y : %f, Z : %q", x, y, z)
// 	fmt.Println()
// 	fmt.Printf("Score : %d ", score)
// 	fmt.Println()
// 	fmt.Printf(" Same Verb X : %v, Y : %v, Z : %v", x, y, z)
// 	fmt.Println()
// 	fmt.Printf("Score : %v \n", score)

// 	fmt.Printf("Upto 4 decimal : %.4f\n", m)

// }
