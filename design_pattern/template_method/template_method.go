package template_method

import "fmt"

type IHummerModel interface {
	Start()
	Stop()
	Alarm()
	EngineBoom()
}

type HummerModel struct {
	model IHummerModel
}

func (hm *HummerModel) Run() {
	hm.model.Start()
	hm.model.EngineBoom()
	hm.model.Alarm()
	hm.model.Stop()
}

type HummerH1Model struct {
	HummerModel
}

func NewHummerH1Model() *HummerModel {
	return &HummerModel{
		&HummerH1Model{},
	}
}

func (h1 *HummerH1Model) Start() {
	fmt.Println("hummer h1 start")
}

func (h1 *HummerH1Model) EngineBoom() {
	fmt.Println("hummer h1 engine boom")
}

func (h1 *HummerH1Model) Alarm() {
	fmt.Println("hummer h1 alarm")
}

func (h1 *HummerH1Model) Stop() {
	fmt.Println("hummer h1 stop")
}

type HummerH2Model struct {
	HummerModel
}

func NewHummerH2Model() *HummerModel {
	return &HummerModel{
		&HummerH2Model{},
	}
}

func (h2 *HummerH2Model) Start() {
	fmt.Println("hummer h2 start")
}

func (h2 *HummerH2Model) EngineBoom() {
	fmt.Println("hummer h2 engine boom")
}

func (h2 *HummerH2Model) Alarm() {
	fmt.Println("hummer h2 alarm")
}

func (h2 *HummerH2Model) Stop() {
	fmt.Println("hummer h2 stop")
}
