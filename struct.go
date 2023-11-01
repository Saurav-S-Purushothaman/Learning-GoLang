package main

import "fmt" 

// we will learn about struct in go

type Person struct{
	name string 
	age int
}
type rect struct{
	width,height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim()  int {
	return 2*r.width + 2*r.height
}

func newPerson(name string) *Person {
	person := &Person{name:name, age:21}
	return person
}

func main(){
    person := newPerson("saurav")
	person.name ="joshua"

	s:= Person{name:"saurav", age:24}
	s_pointer := &s
	fmt.Println("Lets see what value we get", s_pointer)
	fmt.Printf("value of the struct %v\n", person)

	r := rect{width:10, height:5}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	r.width = 20
	fmt.Println(r.area())
	fmt.Println(r.perim())
}