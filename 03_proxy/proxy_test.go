package proxy

import "testing"

func TestNewSchoolGirl(t *testing.T) {
	var jiaojiao = NewSchoolGirl("JiaoJiao")

	var zhuiqiuzhe = NewPursuit(jiaojiao)

	var proxy = NewProxy(zhuiqiuzhe)

	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()

	t.Log("\nend")

}
