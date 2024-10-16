package main

import (
	"fmt"
)

/*
	func main() {
		itemsSold()
	}

	func itemsSold() {
		items := make(map[string]int)
		items["John"] = 41
		items["Celina"] = 109
		items["Micah"] = 24
		for k, v := range items {
			fmt.Printf("%s sold %d items and ", k, v)
			if v < 40 {
				fmt.Println("is below expectations.")
			} else if v > 40 && v <= 100 {
				fmt.Println("meets expectations.")
			} else if v > 100 {
				fmt.Println("exceeded expectations.")
			}
		}
	}
*/

/*
func main() {
	for i := 1; i <= 15; i++ {
		n, s := fizzBuzz(i)
		fmt.Printf("Results: %d %s\n", n, s)
	}
}

func fizzBuzz(i int) (int, string) {
	switch {
	case i%15 == 0:
		return i, "FizzBuzz"
	case i%3 == 0:
		return i, "Fizz"
	case i%5 == 0:
		return i, "Buzz"
	}
	return i, ""
}
*/

// Weekday type definition
type Weekday int

// Enum for days of the week
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Employee type definition
type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

// Developer type definition
type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

// LogHours method for Developer
func (d *Developer) LogHours(day Weekday, hours int) {
	if day >= Sunday && day <= Saturday {
		d.WorkWeek[day] = hours
	}
}

// HoursWorked method for Developer
func (d *Developer) HoursWorked() int {
	totalHours := 0
	for _, hours := range d.WorkWeek {
		totalHours += hours
	}
	return totalHours
}

func main() {
	// Initialize Employee
	emp := Employee{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	// Initialize Developer
	dev := Developer{
		Individual: emp,
		HourlyRate: 50,
		WorkWeek:   [7]int{0, 0, 0, 0, 0, 0, 0},
	}

	// Log hours for Monday and Tuesday
	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)

	// Print the hours for Monday and Tuesday
	fmt.Printf("Hours worked on Monday: %d\n", dev.WorkWeek[Monday])
	fmt.Printf("Hours worked on Tuesday: %d\n", dev.WorkWeek[Tuesday])
	fmt.Println("Hours worked this week: ", dev.HoursWorked())
}
