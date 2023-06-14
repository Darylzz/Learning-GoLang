package main

import "fmt"

// func main() {
// 	var name string = "Cheer Ingkarat"
// 	var age int = 25
// 	var score float32 = 95.8
// 	var isPass bool = true
// 	fmt.Println("My name is", name)
// 	fmt.Println("Age = ",age)
// 	fmt.Println("Score ",score)
// 	fmt.Println("Pass Exam", isPass)
// }

// func main() {
// 	const name string = "Cheer Ingkarat"
// 	fmt.Println("My name is", name)
// }

// func main() {
// 	name:= "Cheer Ingkarat"
// 	age:= 25
// 	score:= 95.8
// 	fmt.Printf("Type name = %T\n", name)
// 	fmt.Printf("Type age = %T\n",age)
// 	fmt.Printf("Type score = %T", score)
// }

//Operator
// func main() {
// 	// var number1 int = 10
// 	// var number2 int = 3
// 	var number1, number2 int = 10,3
// 	fmt.Println("Sum =", number1 + number2)
// 	fmt.Println("Sum =", number1 - number2)
// 	fmt.Println("Sum =", number1 * number2)
// 	fmt.Println("Sum =", number1 / number2)
// 	fmt.Println("Sum =", number1 % number2)
// }

//Check Equal
// func main() {
// 	const number1, number2 int = 10,3
// 	fmt.Println("Equal? =", number1 == number2)
// 	fmt.Println("Not Equal? =", number1 != number2)
// 	fmt.Println(number1, ">", number2, "=", number1 > number2)
// 	fmt.Println(number1, "<", number2, "=", number1 < number2)
// 	fmt.Println(number1, ">=", number2, "=", number1 >= number2)
// }

// Scanf
// func main() {
// 	// var name string
// 	var score int
// 	// fmt.Print("Fill student name =")
// 	// fmt.Scanf("%s", &name)
// 	// fmt.Println("Hello student name =", name)

// 	fmt.Print("Fill score student =")
// 	fmt.Scanf("%d", &score)
// 	fmt.Println("Score student =", score)
// }

// if...else
// func main() {
// 	var number int
// 	fmt.Print("Fill number =")
// 	fmt.Scanf("%d", &number)

// 	if(number == 1) {
// 		fmt.Println("Open book bank")
// 	}else if(number == 2) {
// 		fmt.Println("Deposit money")
// 	}else {
// 		fmt.Println("Invalid Number")
// 	}
// }

// Switch case
// func main() {
// 	var number int
// 	fmt.Print("Fill Number =")
// 	fmt.Scanf("%d", &number)

// 	switch number {
// 	case 1: fmt.Println("Open Book Bank")
// 	case 2: fmt.Println("Deposit Money")
// 	default: fmt.Println("Invalid")
// 	}
// }

// Array
// func main() {
// 	 numbers := [3]int{100,200,300}
// 	 numbers2 := [...]int{400,500,600,700,800,900}
// 	fmt.Println(numbers)
// 	size:= len(numbers)
// 	fmt.Println("Size of array =", size)
// 	fmt.Println("Not fix size in array =", numbers2)
// }

// slice
func main() {
	numbers:= []int{100,200}// default
	numbers = append(numbers, 300)
	fmt.Println(numbers)
	fmt.Println(numbers[:])
	fmt.Println(numbers[1:])
}