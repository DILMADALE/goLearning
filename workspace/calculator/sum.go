package calculator

var LogMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func InternalSum(number int) int {
	return number - 1
}

func Sum(number1, number2 int) int {
	return number1 + number2
}
