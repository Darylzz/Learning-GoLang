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

func main() {
	name:= "Cheer Ingkarat"
	age:= 25
	score:= 95.8
	fmt.Printf("Type name = %T\n", name)
	fmt.Printf("Type age = %T\n",age)
	fmt.Printf("Type score = %T", score)
}