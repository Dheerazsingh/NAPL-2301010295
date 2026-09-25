package main

import (
	"Lab2/mathutil"
	"Lab2/strop"
	"fmt"
)

func main() {

	word := "Reusability"
	fmt.Printf("Original string: %s\n", word)
	fmt.Printf("Reversed string: %s\n", strop.Reverse(word))
	fmt.Printf("Vowel count: %d\n", strop.CountVowels(word))

	n := 5
	base, exp := 2.0, 10

	fmt.Printf("Factorial of %d: %d\n", n, mathutil.Factorial(n))
	fmt.Printf("%.0f^%d = %.0f\n", base, exp, mathutil.Power(base, exp))
}
