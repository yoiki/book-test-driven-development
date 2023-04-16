package model

type Bank struct {
}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Reduce(expression Expression) Money {
	return expression.Reduce()
}
