package factorymethod

import "fmt"

type IFactory interface {
	CreateLeiFeng() LeiFeng
}

//大学生学雷锋工厂
type UndergraduateFactory struct {
}

func (uf *UndergraduateFactory) CreateLeiFeng() LeiFeng {
	return &UndergraduateFactory{}
}

// 社区志愿者工厂
type VolunteerFactory struct {
}

func (vf *VolunteerFactory) CreateLeiFeng() LeiFeng {
	return &VolunteerFactory{}
}

// 雷锋接口
type LeiFeng interface {
	Sweep() // 扫地
	Wash()  // 洗衣
}

func (vf *VolunteerFactory) Sweep() {
	fmt.Println("志愿者学雷锋扫地")
}

func (vf *VolunteerFactory) Wash() {
	fmt.Println("志愿者学雷锋洗衣")
}

func (vf *UndergraduateFactory) Sweep() {
	fmt.Println("大学生学雷锋扫地")
}

func (vf *UndergraduateFactory) Wash() {
	fmt.Println("大学生学雷锋洗衣")
}
