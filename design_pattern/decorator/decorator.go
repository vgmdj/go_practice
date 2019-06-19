package decorator

import (
	"fmt"
	"time"
)

type ILog interface {
	Println(string)
}

type Log struct {
}

func (l *Log) Println(s string) {
	fmt.Printf("%s\n", s)
}

type Decorator struct {
	component ILog
}

func (d *Decorator) SetLog(c ILog) {
	d.component = c
}

func (d *Decorator) Println(s string) {
	if d.component != nil {
		d.component.Println(s)
	}
}

type TimeLog struct {
	Decorator
}

func (tl *TimeLog) Println(s string) {
	fmt.Printf("[time:  %s] ", time.Now().Format("2006-01-02 15:04:05"))
	tl.component.Println(s)

}

type PathLog struct {
	Decorator
}

func (pl *PathLog) Println(s string) {
	fmt.Printf("[path: %s] ", "/path")
	pl.component.Println(s)

}
