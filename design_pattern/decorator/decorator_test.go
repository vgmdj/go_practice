package decorator

import "testing"

func TestDecorator(t *testing.T) {
	l := new(Log)
	tl := new(TimeLog)
	pl := new(PathLog)

	tl.SetLog(l)
	pl.SetLog(tl)

	pl.Println("this is a log test")

}

func TestDecorator2(t *testing.T) {
	p := &Person{Name: "test"}
	ts := new(TShirts)
	bt := new(BigTrouser)

	ts.Decorate(p)
	bt.Decorate(ts)

	bt.Show()

}
