const assert = require('assert');

class Dollar {
    constructor(amount) {
      this.amount = amount;
    }

    times(multiplier) {
        return new Dollar(this.amount * multiplier)
    }
}

fiver = new Dollar(5);
tenner = fiver.times(2);
assert.strictEqual(tenner.amount, 10);
