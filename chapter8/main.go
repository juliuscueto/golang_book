package main

import "fmt"

func one(xPtr *int)  {
	*xPtr = 1
}

func swap(xPtr *int, yPtr *int)  {
	tmp := *xPtr
	*xPtr = *yPtr
	*yPtr = tmp
}

func main() {
	var(
		x = 10
		y = 4
	)
	swap(&x, &y)
	fmt.Println(x, y)
}