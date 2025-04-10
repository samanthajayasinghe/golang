package main

import "fmt"

type Pizza struct {
	slices   int
	toppings []string
	isBaked  bool
}

// Constructors : Get a new Pizza
func NewPizza(slices int, toppings []string) Pizza {
	if toppings == nil {
		toppings = []string{}
	}
	return Pizza{
		slices:   slices,
		toppings: toppings,
	}
}

func (p Pizza) EatSlice() Pizza {
	p.slices = p.slices - 1
	return p
}

func (p Pizza) Bake() {
	p.isBaked = true
}

func (p Pizza) Print() {
	fmt.Printf("slices:%d \ntoppings:%s \n", p.slices, p.toppings)
}

// Pizza implements Bakeable
type Bakeable interface {
	Bake()
}

// this constructor will return a `Bakeable`
// and not a `Pizza`
func NewUnbakedPizza(toppings []string) Bakeable {
	return Pizza{
		slices:   6,
		toppings: toppings,
	}
}

func main() {
	myPizza := NewPizza(6, nil)
	fmt.Print("=======My Pizza========\n")
	myPizza.Print()

	fmt.Print("======My Kid Pizza========\n")
	myKidPizza := NewPizza(8, []string{"pineapple", "tomato"})
	myKidPizza.Print()
}
