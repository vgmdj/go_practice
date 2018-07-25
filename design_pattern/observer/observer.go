package observer

import "fmt"

type Event struct {
	Data string
}

type Subject interface {
	Register(observer Observer)
	Destroy(observer Observer)

	Notify(event Event)
}

type Observer interface {
	Update(event Event)
}

type ConcreteSubject struct {
	Observers map[Observer]struct{}
}

func (cs *ConcreteSubject) Register(observer Observer) {
	if cs.Observers == nil{
		cs.Observers = make(map[Observer]struct{})
	}
	cs.Observers[observer] = struct{}{}
}

func (cs *ConcreteSubject) Destroy(observer Observer) {
	delete(cs.Observers, observer)
}

func (cs *ConcreteSubject) Notify(event Event) {
	for obs, _ := range cs.Observers {
		obs.Update(event)
	}
}

type ConcreteObserver struct {
	ID string
}

func (co *ConcreteObserver) Update(event Event) {
	fmt.Printf("this is observer %s and the event is %s \n", co.ID, event.Data)
}
