package main

import (
	"reflect"
	s "tdd/stocks"
	"testing"
)

var bank s.Bank

func initExchangeRates() {
	bank = s.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}
func TestMultiplication(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := s.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := s.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

func TestAddition(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")
	fifteenDollars := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars, err := portfolio.Evaluate(bank, "USD")

	assertEqual(t, fifteenDollars, *portfolioInDollars)
	assertNil(t, err)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := s.NewMoney(17, "USD")
	actualValue, err := portfolio.Evaluate(bank, "USD")

	assertEqual(t, expectedValue, *actualValue)
	assertNil(t, err)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	elevenHundredWon := s.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)

	expectedValue := s.NewMoney(2200, "KRW")
	actualValue, err := portfolio.Evaluate(bank, "KRW")

	assertEqual(t, expectedValue, *actualValue)
	assertNil(t, err)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	oneEuro := s.NewMoney(1, "EUR")
	oneWon := s.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage :=
		"Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	value, actualError := portfolio.Evaluate(bank, "Kalganid")

	assertNil(t, value)
	assertEqual(t, expectedErrorMessage, actualError.Error())
}

func TestAddTwoMoneysInSameCurrency(t *testing.T) {
	fiveEuros := s.NewMoney(5, "EUR")
	tenEuros := s.NewMoney(10, "EUR")
	expectedValue := s.NewMoney(15, "EUR")

	actualValue := fiveEuros.Add(&tenEuros)
	assertEqual(t, expectedValue, *actualValue)

	actualValue = tenEuros.Add(&fiveEuros)
	assertEqual(t, expectedValue, *actualValue)
}

func TestAddTwoMoneysInDiffernetCurrencies(t *testing.T) {
	euro := s.NewMoney(1, "EUR")
	dollar := s.NewMoney(1, "USD")

	assertNil(t, dollar.Add(&euro))
	assertNil(t, euro.Add(&dollar))
}

func TestConversionWithDifferentRatesBetweenTwoCurrencies(t *testing.T) {
	initExchangeRates()
	tenEuros := s.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, s.NewMoney(12, "USD"), *actualConvertedMoney)

	bank.AddExchangeRate("EUR", "USD", 1.3)
	actualConvertedMoney, err = bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, s.NewMoney(13, "USD"), *actualConvertedMoney)
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "Kalganid")
	assertNil(t, actualConvertedMoney)
	assertEqual(t, "EUR->Kalganid", err.Error())
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Expected to be nil, found %+v", actual)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}
