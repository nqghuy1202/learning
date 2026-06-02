package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   string
	phone     string
	email     string
}

type Address struct {
	street  string
	city    string
	country string
}

func (p *Person) fullName() string {
	fmt.Printf("pointer reciever: %p", p)

	return p.firstName + " " + p.lastName
}

func newPerson(firstName string, lastName string, age int) *Person {
	if age < 0 {
		return nil
	}

	p := new(Person)

	p.firstName = firstName
	p.lastName = lastName
	p.age = age

	return p
}

func buildPerson() *Person {
	return new(Person)
}

func (p *Person) withFirstName(firstName string) *Person {
	p.firstName = firstName
	return p
}

func (p *Person) withLastName(lastName string) *Person {
	p.lastName = lastName
	return p
}

func (p *Person) withAge(age int) *Person {
	p.age = age
	return p
}
