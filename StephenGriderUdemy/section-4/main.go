package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo //just contactInfo also works
}

func main() {
	// newPerson := person{firstName: "abc", lastName: "def"}
	// // newPerson := person{"abc", "def"} //order needs to be maintained
	// fmt.Println(newPerson)
	// fmt.Printf("%+v", newPerson) 
	// //returns
	// //{abc def}
	// // {firstName:abc lastName:def}% 
	
	newPerson := person{
		firstName: "abc", 
		lastName: "def", 
		contact: contactInfo{
			email: "abc@gmail.com", 
			zip: 12345,
			},
	}
	// fmt.Printf("%+v", newPerson)
	// fmt.Println(newPerson.contact.email)
	// newPersonPointer := &newPerson //pointer
	newPerson.updateName("ABC") //not a pointer but still works
	newPerson.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //pass by reference
}

func (p person) print() {
	fmt.Printf("%+v", p)
}