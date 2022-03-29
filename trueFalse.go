package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 88

	//Insert your code from here
	if value < secretValue {
		resultMessge = "Too low, try again next time!"
	} else if value > secretValue {
		resultMessge = "Too high, try again next time!"
	} else {
		resultMessge = "Well Done! Your guess is correct"
	}
	//do not remove this line
	return resultMessge
}

func main() {
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	result := compare(guess)
	print(result)
}
