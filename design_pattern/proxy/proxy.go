package proxy

import "fmt"

type Subject interface {
	Request()
}

type ConcreteSubject struct {
}

func (cs *ConcreteSubject) Request() {
	fmt.Println("this is real request")
}

type Proxy struct {
	s Subject
}

func (p *Proxy) Request() {
	if p.s == nil {
		p.s = new(ConcreteSubject)
	}

	p.Before()
	p.s.Request()
	p.After()
}

func (p *Proxy) Before() {
	fmt.Println("预处理")
}

func (p *Proxy) After() {
	fmt.Println("后置处理")
}
