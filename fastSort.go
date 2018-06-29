package main

import "fmt"

var (
	test  = []int{1,4,2,3,7,5,8}
)
func fastSort(a [] int, left int, right int) []int {
	//when the array's length less or equal 1,just return the array
	if right - left <= 1 {
		return a
	}
	var (
		base int
	)
	base = a[left]
	leftIndex,rightIndex := left,right
	for leftIndex < rightIndex {
		for leftIndex < rightIndex {
			if base < a[rightIndex] {
				a[leftIndex] = a[rightIndex]
				leftIndex ++
				break
			} else {
				rightIndex --
			}
		}

		for leftIndex < rightIndex {
			if base > a[leftIndex] {
				a[rightIndex] = a[leftIndex]
				rightIndex --
				break
			} else {
				leftIndex ++
			}
		}
	}
	a[leftIndex] = base
	fmt.Println(a)
	fastSort(a, left, leftIndex - 1)
	fastSort(a, rightIndex + 1, right)
	return a
}

func main () {
	fastSort(test, 0, len(test) - 1)
	fmt.Println(test)
}
