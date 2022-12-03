package model

type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) *Dollar {
	return &Dollar{amount: d.amount * multiplier}
}

func (d *Dollar) equals(another Dollar) bool {
	return d.amount == another.amount
}
