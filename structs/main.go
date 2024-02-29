package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) toString() string {
	return fmt.Sprintf("%s %s email is %s and phone is %s", p.firstName, p.lastName, p.contactInfo.email, p.contactInfo.phone)
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.contactInfo = contactInfo{email: "alex@gmail.com", phone: "4343243"}
	alex.updateName("Alexe")
	fmt.Println(alex.toString())
	fmt.Printf("%+v", alex)
}
