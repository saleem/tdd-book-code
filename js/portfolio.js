const Money = require('./money');

class Portfolio {
    constructor() {
        this.moneys = [];
    }

    add(...moneys) {
        this.moneys = this.moneys.concat(moneys);
    }

    evaluate(currency) {
        let total = this.moneys.reduce((sum, money) => {
            return sum + money.amount;
        }, 0);
        return new Money(total, currency);
    }
}

module.exports = Portfolio;
