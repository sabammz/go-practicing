package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	Person1 := func(firstname string, lastname string) Person {
		fmt.Println("we are calling anonymous function")
		return Person{
			Firstname: firstname,
			Lastname: lastname,
			Age : 20,
		}
	}
	person2:= person1(firstname,"ali", lastname,"goli")
	fmt.Println(person2.firstname)
}

func (p *Person)GetAge() int{
	return p.Age

}

func GetNewPerson(createNewPerson func(firstname string, lastname string)Person) Person{
	return createNewPerson(firstame "ali" , lastname "goli")
}