const assert = require('assert');

class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency;
    }

    times(multiplier) {
        return new Money(this.amount * multiplier, this.currency);
    }

    divide(divisor) {
        return new Money(this.amount / divisor, this.currency);
    }
}

fiveDollars = new Money(5, "USD");
tenDollars = new Money(10, "USD");
assert.deepStrictEqual(fiveDollars.times(2), tenDollars);

tenEuros = new Money(10, "EUR");
twentyEuros = new Money(20, "EUR");
assert.deepStrictEqual(tenEuros.times(2), twentyEuros);

originalMoney = new Money(4002, "KRW")
expectedMoneyAfterDivision = new Money(1000.5, "KRW")
assert.deepStrictEqual(originalMoney.divide(4), expectedMoneyAfterDivision)
