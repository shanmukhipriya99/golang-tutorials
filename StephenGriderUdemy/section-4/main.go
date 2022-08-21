package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	newPerson := person{firstName: "abc", lastName: "def"}
	// newPerson := person{"abc", "def"} //order needs to be maintained
	fmt.Println(newPerson)
	fmt.Printf("%+v", newPerson) 
	//returns
	//{abc def}
	// {firstName:abc lastName:def}%                          
}