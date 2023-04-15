package model

import (
	"errors"
	"fmt"
)

type MonetaryUnit int

const (
	MonetaryUnitUnknown MonetaryUnit = iota
	MonetaryUnitUSD
	MonetaryUnitCHF
)

var moneyUnit = map[MonetaryUnit]string{
	MonetaryUnitUSD: "USD",
	MonetaryUnitCHF: "CHF",
}

func NewMoneyUnit(m string) (MonetaryUnit, error) {
	u := MonetaryUnitUnknown
	for k, v := range moneyUnit {
		if v == m {
			u = k
			break
		}
	}
	if err := u.validation(); err != nil {
		return MonetaryUnitUnknown, errors.New(fmt.Sprintf("Invalid monetary unit: %s\n", m))
	}
	return u, nil
}

func (m MonetaryUnit) validation() error {
	_, ok := moneyUnit[m]
	if !ok && m == MonetaryUnitUnknown {
		return errors.New("Invalid monetary unit\n")
	}
	return nil
}

func (m MonetaryUnit) String() (string, error) {
	u, ok := moneyUnit[m]
	if !ok {
		return "Unknown", errors.New("Invalid monetary unit\n")
	}
	return u, nil
}
