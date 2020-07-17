package proxy

import "fmt"

// 送礼物
type GiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

// 被追求者
type SchoolGirl struct {
	name string
}

func NewSchoolGirl(name string) *SchoolGirl {
	return &SchoolGirl{name: name}
}

//追求者
type Pursuit struct {
	mm SchoolGirl
}

func NewPursuit(mm *SchoolGirl) *Pursuit {
	return &Pursuit{mm: *mm}
}

func (p *Pursuit) GiveDolls() {
	fmt.Printf("\n%s 送你洋娃娃", p.mm.name)
}

func (p *Pursuit) GiveFlowers() {
	fmt.Printf("\n%s 送你鲜花", p.mm.name)
}

func (p *Pursuit) GiveChocolate() {
	fmt.Printf("\n%s 送你巧克力", p.mm.name)
}

// 代理
type Proxy struct {
	gg Pursuit
}

func NewProxy(gg *Pursuit) *Proxy {
	return &Proxy{gg: *gg}
}

func (p *Proxy) GiveDolls() {
	p.gg.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.gg.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.gg.GiveChocolate()
}
