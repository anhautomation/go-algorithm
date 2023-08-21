package main

import(
	"time"
	"fmt"
	"go-algorithm/algorithms"
)

func main() {

	firstArr := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	firstStart := time.Now()
	isFound := algorithms.LinerSearch(firstArr,89)
	firstTimeElapsed := time.Since(firstStart)
	fmt.Println("linear search for 89:", isFound, "Time elapsed:", firstTimeElapsed)

	secondArr := []int{89, 55, 34, 21, 13, 8, 5, 3, 2, 1, 1, 144, 233, 377, 610}
	secondStart := time.Now()
	isFound = algorithms.LinerSearch(secondArr, 612)
	secondTimeElapsed := time.Since(secondStart)
	fmt.Println("linear search for 612:", isFound, "Time elapsed:", secondTimeElapsed)

	bubbleSortArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", bubbleSortArr)
	algorithms.BubbleSort(bubbleSortArr)
	fmt.Println("Sorted array:", bubbleSortArr)

	binarySearchArr := []int{11, 12, 22, 25, 34, 64, 90}
	target := 25
	index := algorithms.BinarySearch(binarySearchArr, target)
	if index != -1 {
		fmt.Println("Element", target, "found at index ", index)
	} else {
		fmt.Println("Element ", target," not found")
	}
	
	n := 10
	fmt.Println("Fibonacci of ", n," is ", algorithms.Fibonacci(n))

	fmt.Println("factorial of 5 is ", algorithms.Factorial(5))

	solutions := algorithms.SolveNQueens(4)
	for _, solution := range solutions {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println()
	}
	
}