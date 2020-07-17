package decorator

import "fmt"

// Component
type Person interface {
	Show()
}

// ConcreteComponent
type ConcretePerson struct {
	name string
}

func NewConcretePerson(name string) *ConcretePerson {
	return &ConcretePerson{name: name}
}

func (p *ConcretePerson) Show() {
	fmt.Printf("装扮的%s", p.name)
}

//Decorator
type Finery interface {
	Person
}

type TShirt struct {
	Finery
}

func NewTShirt(finery Finery) *TShirt {
	return &TShirt{Finery: finery}
}

func (ts *TShirt) Show() {
	fmt.Println("T恤 ")
	ts.Finery.Show()
}

type BigTrouser struct {
	Finery
}

func NewBigTrouser(finery Finery) *BigTrouser {
	return &BigTrouser{Finery: finery}
}

func (bt *BigTrouser) Show() {
	fmt.Println("大裤衩 ")
	bt.Finery.Show()
}
