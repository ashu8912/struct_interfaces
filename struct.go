package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int64
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (p *person) updateName(name string) {
	p.firstName = name
}
func main() {
	ashu := person{
		firstName: "ashu",
		lastName:  "ghildiyal",
		contactInfo: contactInfo{
			email:   "ashughildiyal5@gmail.com",
			zipCode: 110084,
		},
	}

	ashu.updateName("ajay")
	ashu.print()

}
