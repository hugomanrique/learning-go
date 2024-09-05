package animal

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

func (p *Dog) Sound() {
	fmt.Println(p.Name, "Hace: ", "Woof")
}

type Cat struct {
	Name string
}

func (p *Cat) Sound() {
	fmt.Println(p.Name, "Hace: ", "Meow")
}

func DoSound(animal Animal) {
	animal.Sound()
}
