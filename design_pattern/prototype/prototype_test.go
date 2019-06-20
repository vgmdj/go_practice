package prototype

import (
	"testing"
)

func TestProtoType(t *testing.T) {
	resume := NewResume("test", "man", 18)
	resume.SetEducation("哈工大", "哈尔滨")
	resume.AddWorkExperience("华为", "2016年", "2017年")

	resume2 := resume.Clone()
	resume2.SetEducation("清华", "北京")
	resume2.AddWorkExperience("阿里", "2017年", "2018年")

	t.Log(resume)
	t.Log(resume2)

}
