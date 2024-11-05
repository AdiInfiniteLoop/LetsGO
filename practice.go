package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1
	for i := 1; i < 11; i++ {
		fmt.Printf("%v ", i)
	}
	// 2
	var n int
	fmt.Println("Enter the value of n")
	fmt.Scan(&n)
	var sumofN int = 0
	for i := 1; i <= n; i++ {
		sumofN = i + sumofN
	}
	fmt.Println(sumofN)
	// 3.Print even numbers
	num := 1
	for num <= 50 {
		if num%2 == 0 {
			fmt.Println(num)
		}
		num = num + 1
	}

	// 4. linear search in Go
	// var arr [10]int; or
	// arr := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	key := 2
	arr := [10]int{5, 6, 43, 2, 54, 6, 7, 87, 9, 2}
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			fmt.Println(i-1, "th index")
		}
	}
	/*--------------------------SLICES-----------------------------*/
	nums := []string{} // If using this syntax of var x = []string{}, {} is a must
	// var x = []string{1,3,4,43,2};// just do not input the size
	nums = append(nums, "new name")

	// Delete or Slice the slices
	// nums := append(nums[:1]);//to delete first element
	// nums := append(nums[3:])

	newarr := make([]int, 4)
	newarr[0] = 10
	newarr[1] = 20
	newarr[2] = 30
	newarr[3] = 40
	// newarr[4] = 50; this line causes error but
	newarr = append(newarr, 50)
	newarr = append(newarr, 60)
	newarr = append(newarr, 70)

	fmt.Println(newarr)
	sort.Ints(newarr)
	fmt.Println(newarr)
	sort.IntsAreSorted(newarr)

	// Remove a value from a slice
	slice := []int{1, 3, 5, 6, 4, 1}
	// to delete element at slice
	idx := 2
	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Println(slice)
}
