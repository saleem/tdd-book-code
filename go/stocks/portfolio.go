package stocks

import "errors"

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	p = append(p, money)
	return p
}

func (p Portfolio) Evaluate(currency string) (Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, m := range p {
		if convertedAmount, ok := convert(m, currency); ok {
			total = total + convertedAmount
		} else {
			failedConversions = append(failedConversions,
				m.currency+"->"+currency)
		}
	}
	if len(failedConversions) == 0 {
		return NewMoney(total, currency), nil
	}
	failures := "["
	for _, f := range failedConversions {
		failures = failures + f + ","
	}
	failures = failures + "]"
	return NewMoney(0, ""),
		errors.New("Missing exchange rate(s):" + failures)
}

func convert(money Money, currency string) (float64, bool) {
	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	if money.currency == currency {
		return money.amount, true
	}
	key := money.currency + "->" + currency
	rate, ok := exchangeRates[key]
	return money.amount * rate, ok
}
