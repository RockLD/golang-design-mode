package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	var p Person = &ConcretePerson{
		name: "RockLd",
	}
	p = NewBigTrouser(p)
	p = NewTShirt(p)
	p.Show()
}
