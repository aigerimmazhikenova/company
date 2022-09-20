package main

import "fmt"

type Person struct {
	name string
}

func newPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

func (p *Person) handleEvent(v []string) {
	fmt.Printf("Hi dear %s \n Vacancies updated: ", p.name)
	for index, value := range v {
		fmt.Printf("%v) %v\n", index+1, value)
	}
}
