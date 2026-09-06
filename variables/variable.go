package main

import "fmt"

// A sample of explicit variable declaration
var  name = "tobi"


func main() {
	printNameAndAge()
	printTwoNumbers()
}

func printNameAndAge() {
	age, name := 27, "tobi"
	age = age + 3
	fmt.Printf("my name is %s and i am %d years old\n", name, age)
}

func printTwoNumbers() {
	num1 := 30
	num2 := 40
	fmt.Println(num1 * num2)
}


// practice explicit declaration using var
// implicit declaration using :=
// practice writing programs that use variables
