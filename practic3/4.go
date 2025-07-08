package main

import "fmt"

type Person struct {
	age  int
	name string
}

func (p *Person) info() {
	fmt.Println(p.age, p.name)
}

type Employee struct {
	Person
	EmpId int
}

func (e *Employee) info() {
	fmt.Println(e.EmpId, e.age, e.name)
}

func main() {
	a := Person{age: 5, name: "Vik"}
	b := Employee{Person: a, EmpId: 11}
	b.info()

}
