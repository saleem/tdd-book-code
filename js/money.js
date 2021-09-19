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

    add(money) {
        if (money.currency !== this.currency) {
            throw new Error("Cannot add " + money.currency + " to " + this.currency);
        }
        return new Money(this.amount + money.amount, this.currency);
    }

}

module.exports = Money;
