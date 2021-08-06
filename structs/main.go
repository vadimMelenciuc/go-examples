package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Alex"

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 123,
		},
	}

	fmt.Printf("%+v\n", jim)

	sally := person{"Sally", "Person", contactInfo{"sally@gmail.com", 1234}}
	fmt.Printf("%+v\n", sally)

	fmt.Println("_______________________________________")
	alex.print()
	jim.print()
	sally.print()

	fmt.Println("_______________________________________")
	jim.newFirstNameFail("Jeremias") // Doesn't updates original name
	jim.print()

	fmt.Println("_______________________________________")
	jimPointer := &jim
	jimPointer.newFirstName("Jeremias")
	jim.print()

	fmt.Println("_______________________________________")
	(&jim).newFirstName("Jeremias 2")
	jim.print()

	fmt.Println("_______________________________________")
	// This will work too, go will convert the type to a pointer
	jim.newFirstName("Jeremias 3")
	jim.print()

}

func (p person) newFirstNameFail(newName string) {
	p.firstName = newName
}

func (pointerToPerson *person) newFirstName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {

	fmt.Printf("%+v\n", p)
}
