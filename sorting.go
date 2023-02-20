package main

import (
	"fmt"
)

// Bubble sort
func bubble(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

// Selection Sort
func Selection(arr []int) {
	for i := 0; i < len(arr); i++ {
		small := i
		for j := i + 1; j < len(arr); j++ {
			if arr[small] > arr[j] {
				small = j
			}
		}
		if small != i {
			temp := arr[i]
			arr[i] = arr[small]
			arr[small] = temp
		}
	}
}

// Insertion sort
func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && temp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}

}

// Merge sort

func Merge(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	return Sort(Merge(left), Merge(right))
}

func Sort(left, right []int) []int {
	size := len(left) + len(right)
	sorted := make([]int, size)
	i := 0
	j := 0
	k := 0
	for len(left) > i && len(right) > j { // for i,j = 1
		if left[i] > right[j] {
			sorted[k] = right[j]
			j++
			k++
		} else {
			sorted[k] = left[i]
			i++
			k++
		}
	}
	for len(left) > i {
		sorted[k] = left[i]
		k++
		i++
	}
	for len(right) > j {
		sorted[k] = right[j]
		k++
		j++
	}
	return sorted
}

// Quick sort
func Quick(arr []int) []int {
	HelpQuick(arr, 0, len(arr)-1)
	return arr
}
func HelpQuick(arr []int, start, end int) []int {
	if start >= end {
		return arr
	}
	prvt := start
	left := start + 1
	right := end
	for right >= left {
		if arr[prvt] < arr[left] && arr[prvt] > arr[right] {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
			left++
			right--
		}
		if arr[prvt] > arr[left] {
			left++
		}
		if arr[prvt] < arr[right] {
			right--
		}
	}
	Swap(arr, prvt, right)
	HelpQuick(arr, start, right)
	HelpQuick(arr, right+1, end)
	return arr
}
func Swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Array printing
func Print(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}
func main() {
	arr := []int{5, 9, 8, 1, 3, 6, 4, 7}
	fmt.Println("First array is: ", arr)
	v := Quick(arr)
	fmt.Print(v)
	// Print(new)
	// Selection(arr)
	// fmt.Println(arr)
}
