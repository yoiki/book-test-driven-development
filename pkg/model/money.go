package model

import (
	"reflect"
)

type Expression interface {
	Reduce() Money
}

type Money struct {
	amount       int
	monetaryUnit MonetaryUnit
}

func NewMoney(amount int, unit MonetaryUnit) Money {
	m := Money{amount: amount, monetaryUnit: unit}
	return m
}

func (m Money) Reduce() Money {
	return m
}

func (m Money) Amount() int {
	return m.amount
}

func (m Money) MonetaryUnit() MonetaryUnit {
	return m.monetaryUnit
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, monetaryUnit: m.monetaryUnit}
}

func (m Money) Equals(another Money) bool {
	return reflect.DeepEqual(m, another)
}

func (m Money) Plus(another Money, unit MonetaryUnit) Sum {
	return NewSum(m, another, unit)
}

type Sum struct {
	augend Money
	addend Money
	unit   MonetaryUnit
}

func NewSum(augend Money, addend Money, unit MonetaryUnit) Sum {
	return Sum{augend: augend, addend: addend, unit: unit}
}

func (s Sum) Reduce() Money {
	return NewMoney(s.augend.Amount()+s.addend.Amount(), s.unit)
}
