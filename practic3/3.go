package main

import "fmt"

func main() {
	something := make([]int, 8, 10) //first int is len, second is capacity
	something = append(something, 2)
	fmt.Println(something, len(something), cap(something))
}
