package main

import "fmt"

func main() {

	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)

	var arr = make([]int, size)

	for i:=0; i<size; i++ {
		fmt.Printf("Enter %dth element: ", i+1)
		fmt.Scan(&arr[i])
	}
	
	var result= reverseArray(arr);
	
	fmt.Println("Your Initial array is: ", arr)
	fmt.Println("Your reversed array is: ", result)
}

func reverseArray(arr []int) []int {
	
	for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
