package simplefactory

import (
	"fmt"
	"testing"
)

func TestCreateOperate(t *testing.T) {
	o1 := CreateOperate("+")
	fmt.Println(o1.GetResult(10, 5))
	//t.Log(o1.GetResult(10,5))

	o2 := CreateOperate("-")
	fmt.Println(o2.GetResult(6, 3))

	o3 := CreateOperate("*")
	fmt.Println(o3.GetResult(6, 3))

	o4 := CreateOperate("/")
	fmt.Println(o4.GetResult(6, 3))

	o5 := CreateOperate("/")
	fmt.Println(o5.GetResult(6, 0))
}
