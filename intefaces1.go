package main

/*
import "fmt"

type Byer interface {
	Bye()
}

type Greeter interface {
	Greet()
}

//enveviendo interfaces
type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s\n", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adios soy %s\n", p.Name)
}

type Text string

func (t Text) Bye() {
	fmt.Printf("adios soy %s\n", t)
}
func (t Text) Greet() {
	fmt.Printf("hola soy %s\n", t)
}

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t valor: %v Tipo: %T\n", g, g)
	}
}

func All(gb ...GreeterByer) {
	for _, g := range gb {
		g.Bye()
		g.Greet()
	}
}

/*
func main() {
	//var g Greeter = &Person{Name: "alejandro"}
	p := Person{Name: "santiago"}

	var t Text = "dayse"
	All(p, t)

}
*/
