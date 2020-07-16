package simplefactory

type Operation interface {
	GetResult(NumberA, NumberB int) int
}

type OperationAdd struct{}

func (o *OperationAdd) GetResult(NumberA, NumberB int) int {
	return NumberA + NumberB
}

type OperationSub struct{}

func (o *OperationSub) GetResult(NumberA, NumberB int) int {
	return NumberA - NumberB
}

type OperationMul struct{}

func (o *OperationMul) GetResult(NumberA, NumberB int) int {
	return NumberA * NumberB
}

type OperationDiv struct{}

func (o *OperationDiv) GetResult(NumberA, NumberB int) int {
	if 0 == NumberB {
		panic("除数不能为0")
	}
	return NumberA / NumberB
}

func CreateOperate(oprate string) Operation {
	switch oprate {
	case "+":
		return &OperationAdd{}
	case "-":
		return &OperationSub{}
	case "*":
		return &OperationMul{}
	case "/":
		return &OperationDiv{}
	}
	return nil
}
