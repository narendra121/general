package main

import "fmt"

type myInterface interface {
	AddUser(m myStruct)
	GetUser() *myStruct
}

type myStruct struct {
	Name  string
	EmpId int
}

func (a *myStruct) AddUser(m myStruct) {
	a.EmpId = m.EmpId
	a.Name = m.Name
}
func (a *myStruct) GetUser() *myStruct {
	return a
}
func main() {
	var myInt myInterface
	var myStr myStruct
	myInt = &myStr
	myInt.AddUser(myStruct{
		"narendra",
		123,
	})
	fmt.Println(myInt.GetUser())
}
