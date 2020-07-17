package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	var p Person = NewConcretePerson("Rock")
	p = NewBigTrouser(p)
	p = NewTShirt(p)
	p.Show()
}
