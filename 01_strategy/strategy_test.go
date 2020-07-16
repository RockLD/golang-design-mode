package strategy

import (
	"fmt"
	"testing"
)

func TestContext_ContextInterface(t *testing.T) {
	c1 := CreateCash("正常收费")
	if nil != c1 {
		fmt.Println(c1.GetResult(100))
	}

	c2 := CreateCash("满300减100")
	if nil != c2 {
		fmt.Println(c2.GetResult(1000))
	}

	c3 := CreateCash("打8折")
	if nil != c3 {
		fmt.Println(c3.GetResult(500))
	}
	println("end")
}
