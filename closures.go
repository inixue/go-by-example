package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextI := intSeq()
	fmt.Println(nextI())
	fmt.Println(nextI())
	fmt.Println(nextI())
	fmt.Println(nextI())

	nextII := intSeq()
	fmt.Println(nextII())
	fmt.Println(nextII())
}
