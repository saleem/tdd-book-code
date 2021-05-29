const Money = require('./money');
const Bank = require('./bank');

class Portfolio {
    constructor() {
        this.moneys = [];
    }
    add() {
        this.moneys = this.moneys.concat(Array.prototype.slice.call(arguments));
    }

    evaluate(bank, currency) {
        let failures = [];
        let sum = 0;
        let total = this.moneys.reduce( (sum, money) => {
            try {
                var convertedMoney = bank.convert(money, currency);
                return sum + convertedMoney.amount;
            }
            catch (error) {
                failures.push(error.message);
                return sum;
            }
          }, 0);

        if (failures.length == 0) {
            return new Money(total, currency);
        }
        throw new Error("Missing exchange rate(s):[" + failures.join() + "]");
    }
}

module.exports = Portfolio;
