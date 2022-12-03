package model

import "reflect"

type Money struct {
	amount       int
	monetaryUnit string
}

func NewDollar(amount int) *Money {
	return &Money{amount: amount, monetaryUnit: "USD"}
}

func NewFranc(amount int) *Money {
	return &Money{amount: amount, monetaryUnit: "CHF"}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) MonetaryUnit() string {
	return m.monetaryUnit
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{amount: m.amount * multiplier, monetaryUnit: m.monetaryUnit}
}

func (m *Money) Equals(another *Money) bool {
	return reflect.DeepEqual(m, another)
}
