package observer

import "testing"

func TestObservers(t *testing.T) {
	ob1 := &ConcreteObserver{ID: "ob1"}
	ob2 := &ConcreteObserver{ID: "ob2"}

	sb := ConcreteSubject{}
	sb.Register(ob1)
	sb.Register(ob2)

	event1 := Event{Data: "event1"}
	sb.Notify(event1)

	sb.Destroy(ob1)

	event2 := Event{Data: "event2"}
	sb.Notify(event2)
}
