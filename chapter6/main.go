package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallestNum := 1000
	for _, value := range x {
		if value < smallestNum {
			smallestNum = value
		}
	}
	fmt.Println("Smallest Number in:")
	fmt.Println(x)
	fmt.Println("is: ", smallestNum)
}