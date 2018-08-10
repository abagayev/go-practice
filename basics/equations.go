package main

/**
 * This is an example of how to solve quadratic equation without math package
 */

import "fmt"

// just another pow function
func pow(n float64, p int) float64 {
	r := n
	for i := 1; i < p; i++ {
		r *= n
	}

	return r
}

// just a shortcut for pow^2
func square(n float64) float64 {
	return pow(n, 2)
}

// this function is going to be used in sqrt
func CountDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}

	return count
}

// calculating in babylonian method
// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method
func sqrt(s float64) float64 {
	x := float64(CountDigits(int(s)))

	for i := 0; i < 10; i++ {
		x = (x + s/x) / 2
	}

	return x
}

// solve quadratic equation and return both x1 and x2
// https://en.wikipedia.org/wiki/Quadratic_formula
func equation(a float64, b float64, c float64) (float64, float64) {
	d := square(b) - 4*a*c
	if d < 0 {
		panic("complex roots, i am not going to solve it")
	}

	x1 := (-b + sqrt(d)) / (2 * a)
	x2 := (-b - sqrt(d)) / (2 * a)

	return x1, x2
}

func main() {
	fmt.Println(equation(1, -1, -2))
}
