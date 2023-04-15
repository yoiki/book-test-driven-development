package model

import (
	"reflect"
)

type Money struct {
	amount       int
	monetaryUnit MonetaryUnit
}

func NewMoney(amount int, unit MonetaryUnit) (Money, error) {
	m := Money{amount: amount, monetaryUnit: unit}
	if err := m.validation(); err != nil {
		return Money{}, err
	}
	return m, nil
}

func (m Money) validation() error {
	if err := m.monetaryUnit.validation(); err != nil {
		return err
	}
	return nil
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

func (m Money) Plus(another Money) Plus {
	//return Money{m.amount + another.amount, m.MonetaryUnit()}
	return Plus{}
}

type Expression interface {
	Aggregate() Money
}

type Plus struct {
}

func (s Plus) Aggregate() Money {
	return Money{}
}
