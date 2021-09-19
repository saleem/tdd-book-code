const Money = require('./money');

class Portfolio {
    constructor() {
        this.moneys = [];
    }

    add(...moneys) {
        this.moneys = this.moneys.concat(moneys);
    }

    evaluate(bank, currency) {
        let failures = [];
        let total = this.moneys.reduce((sum, money) => {
            try {
                let convertedMoney = bank.convert(money, currency);
                return sum.add(convertedMoney);
            }
            catch (error) {
                failures.push(error.message);
                return sum;
            }
        }, new Money(0, currency));
        if (!failures.length) {
            return total;
        }
        throw new Error("Missing exchange rate(s):[" + failures.join() + "]");
    }
}

module.exports = Portfolio;
