const Money = require('./money');

class Portfolio {
    constructor() {
        this.moneys = [];
    }
    add() {
        this.moneys = this.moneys.concat(Array.prototype.slice.call(arguments));
    }

    evaluate(currency) {
        var total = this.moneys.reduce((sum, money) => {
            return sum + this.convert(money, currency);
        }, 0);
        return new Money(total, currency);
    }

    convert(money, currency) {
        var eurToUsd = 1.2;
        var exchangeRates = new Map(); // <1>
        exchangeRates.set("EUR->USD", 1.2);
        exchangeRates.set("USD->KRW", 1100);
        if (money.currency == currency) {
            return money.amount;
        }
        let key = money.currency + "->" + currency;
        return money.amount * exchangeRates.get(key);
    }
}

module.exports = Portfolio;
