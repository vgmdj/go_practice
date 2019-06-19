package decorator

import "fmt"

type IPerson interface {
	Show()
}

type Person struct {
	Name string
}

func (p *Person) Show() {
	fmt.Printf("装扮的%s\n", p.Name)
}

type Finery struct {
	component IPerson
}

func (f *Finery) Decorate(c IPerson) {
	f.component = c
}

func (f *Finery) Show() {
	if f.component != nil {
		f.component.Show()
	}
}

type TShirts struct {
	Finery
}

func (t *TShirts) Show() {
	fmt.Printf("大T恤 ")
	t.Finery.Show()
}

type BigTrouser struct {
	Finery
}

func (b *BigTrouser) Show() {
	fmt.Printf("垮裤 ")
	b.Finery.Show()
}
