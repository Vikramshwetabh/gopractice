package main

import "fmt"

func Asum() func() int {
	var sum = 0
	return func() int {
		sum += 1
		return sum
	}
}
func main() {
	b := Asum()
	fmt.Println(b())
	fmt.Println(b())

}
