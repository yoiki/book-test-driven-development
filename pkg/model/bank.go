package model

type Bank struct {
}

func NewBank() (Bank, error) {
	b := Bank{}
	if err := b.validation(); err != nil {
		return Bank{}, err
	}
	return Bank{}, nil
}

func (b Bank) validation() error {
	return nil
}

func (b Bank) Reduce(expression Expression) Money {
	return Money{10, MonetaryUnitUSD}
}
