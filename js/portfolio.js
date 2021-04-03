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
        if (money.currency == currency) {
            return money.amount;
        }
        return money.amount * eurToUsd;
    }
}

module.exports = Portfolio;
