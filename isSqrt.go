package main

import  "fmt"

func isSqrt (num int) bool {
	flag := false
	for i := 1;i <= num; i++ {
		if i*i == num {
			flag = true
			break
		}
	}
	return flag
}

func main() {
	test := 4
	fmt.Println((isSqrt(test)))
	test = 36
	fmt.Println(isSqrt(test))
	test = 10000
	fmt.Println(isSqrt(test))
}
