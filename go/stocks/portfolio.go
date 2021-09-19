package stocks

import "errors"

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	p = append(p, money)
	return p
}

func (p Portfolio) Evaluate(bank Bank, currency string) (*Money, error) {
	totalMoney := NewMoney(0, currency)
	failedConversions := make([]string, 0)
	for _, m := range p {
		if convertedMoney, err := bank.Convert(m, currency); err == nil {
			totalMoney = *totalMoney.Add(convertedMoney)
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}
	if len(failedConversions) == 0 {
		return &totalMoney, nil
	}
	failures := "["
	for _, f := range failedConversions {
		failures = failures + f + ","
	}
	failures = failures + "]"
	return nil, errors.New("Missing exchange rate(s):" + failures)
}
