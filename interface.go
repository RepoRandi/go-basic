package main

type HashName interface {
	GetName() string
}
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func sayHelo(hashName HashName) {
	println("Hello " + hashName.GetName())
}

func (person *Person) GetName() string {
	return person.Name
}

func (animal *Animal) GetName() string {
	return animal.Name
}

func main37() {
	person := Person{Name: "John"}
	sayHelo(&person)

	animal := Animal{Name: "Dog"}
	sayHelo(&animal)
}
