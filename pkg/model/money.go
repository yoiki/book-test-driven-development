package model

import "reflect"

type MonetaryUnit int

const (
	Dollar MonetaryUnit = iota
	Franc
)

type Money struct {
	amount       int
	monetaryUnit MonetaryUnit
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) MonetaryUnit() MonetaryUnit {
	return m.monetaryUnit
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{amount: m.amount * multiplier, monetaryUnit: m.monetaryUnit}
}

func (m *Money) Equals(another *Money) bool {
	return reflect.DeepEqual(m, another)
}
