package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
	
}

func factorial(x uint) uint {
	if x==0 {
		return 1
	}
	return x * factorial(x-1)
}

func isHalfEven(x int) (half int, isEven bool) {
	half = x/2
	isEven = x%2 == 0
	return
}

func findLargest(xs ...float64) (ret float64) {
	ret = -1.
	for _, x := range(xs) {
		if ret < x {
			ret = x
		}
	}
	return
}

func fib(i uint) (ret uint) {
	if i==0 || i==1{
		ret = i
	} else {
		ret = fib(i-1) + fib(i-2)
	}
	return
}

func main() {
	fmt.Println(fib(10))
}