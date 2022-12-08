package _structs

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Bob" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func Interface() {
	//p := Person{"Mike"}
	var p Human = &Person{"Mike"}

	DriveCar(p)
}
