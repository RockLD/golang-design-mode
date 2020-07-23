package factorymethod

import "testing"

func TestLeiFent(t *testing.T) {
	// 大学生学雷锋
	factory := &UndergraduateFactory{}
	student := factory.CreateLeiFeng()
	student.Wash()
	student.Sweep()

	// 社区志愿者学雷锋
	factoryA := &VolunteerFactory{}
	volunteer := factoryA.CreateLeiFeng()
	volunteer.Sweep()
	volunteer.Wash()

	t.Logf("end")
}
