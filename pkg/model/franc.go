package model

type Franc struct {
	amount int
}

func (f *Franc) times(multiplier int) *Franc {
	return &Franc{amount: f.amount * multiplier}
}

func (f *Franc) equals(another Franc) bool {
	return f.amount == another.amount
}
