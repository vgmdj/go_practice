package builder

type Builder interface {
	BuildHead()
	BuildBody()
	BuildArms()
	BuildLegs()
}

type Director struct {
	b Builder
}

func (d *Director) SetBuilder(b Builder) {
	d.b = b
}

func (d *Director) Build() {
	d.b.BuildHead()
	d.b.BuildBody()
	d.b.BuildArms()
	d.b.BuildLegs()
}

type ThinBuilder struct {
}

func (tb *ThinBuilder) BuildHead() {

}

func (tb *ThinBuilder) BuildBody() {

}

func (tb *ThinBuilder) BuildArms() {

}

func (tb *ThinBuilder) BuildLegs() {

}
