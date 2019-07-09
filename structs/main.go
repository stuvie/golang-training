package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
}

type person2 struct {
	firstName string
	lastName  string
	contact   contactInfo // field name "contact" can be omitted
}

func main() {
	alex := person{"Alex", "Anderson"}                        // assumed order of fields
	alex2 := person{firstName: "Alex2", lastName: "Anderson"} // more precise

	var skippy person // go assigns zero value to the fields of the struct
	fmt.Println(alex, alex2, skippy)
	fmt.Printf("%+v\n", skippy)

	skippy.firstName = "Alex"
	skippy.lastName = "Morrison"

	fmt.Printf("%+v\n", skippy)

	jim := person2{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPtr := &jim
	jimPtr.updateName("Jimmy")
	jim.print()

	jim.updateName("James") // shortcut way of passing pointer
	jim.print()

	mySlice := []string{"Hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "Bill"
	fmt.Println(*&name)

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func (ptr *person2) updateName(newFirstName string) {
	(*ptr).firstName = newFirstName
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func (p person2) print() {
	fmt.Printf("%+v\n", p)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
