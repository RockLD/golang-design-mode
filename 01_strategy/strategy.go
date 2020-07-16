package strategy

import "math"

// CashSuper CashNormal CashRebate CashReturn CashContext
// 抽象算法
type CashSuper interface {
	AcceptCash(money float64) float64
}

// 正常收费
type CashNormal struct{}

func NewCashNormal() *CashNormal {
	return &CashNormal{}
}

func (n *CashNormal) AcceptCash(money float64) float64 {
	return money
}

// 打折收费
type CashRebate struct {
	moneyRebate float64
}

func NewCashRebate(moneyRebate float64) *CashRebate {
	return &CashRebate{moneyRebate: moneyRebate}
}

func (r *CashRebate) AcceptCash(money float64) float64 {
	return money * r.moneyRebate
}

// 返利收费子类
type CashReturn struct {
	moneyCondition float64
	moneyReturn    float64
}

func NewCashReturn(moneyCondition, moneyReturn float64) *CashReturn {
	return &CashReturn{moneyCondition: moneyCondition, moneyReturn: moneyReturn}
}

func (r *CashReturn) AcceptCash(money float64) float64 {
	result := money
	if money > r.moneyCondition {
		result = money - math.Floor(money/r.moneyCondition)*r.moneyReturn
	}
	return result
}

type Context struct {
	cashSuper CashSuper
}

func CreateCash(stype string) *Context {
	switch stype {
	case "正常收费":
		return &Context{cashSuper: NewCashNormal()}
	case "满300减100":
		return &Context{cashSuper: NewCashReturn(300, 100)}
	case "打8折":
		return &Context{cashSuper: NewCashRebate(0.8)}
	}
	return nil
}

func (c *Context) GetResult(money float64) float64 {
	return c.cashSuper.AcceptCash(money)
}
