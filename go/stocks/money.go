package stocks

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) Money {
	return Money{amount, currency}
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}

func (m Money) Add(other *Money) *Money {
	var result Money
	if m.currency == other.currency {
		result = Money{amount: m.amount + other.amount, currency: m.currency}
		return &result
	}
	return nil
}
