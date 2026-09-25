package lab01

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You are a minor")
	} else {
		fmt.Println("You are an adult")
	}
}
