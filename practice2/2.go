package main

import "fmt"

func main() {
	a := [3]string{"a", "b", "c"}
	b := a[0:1]
	a[0] = "golang"
	fmt.Println(b)
}
