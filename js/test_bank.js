const assert = require('assert');
const Bank = require('./bank');

class BankTest {
  constructor() {
    this.bank = new Bank();
  }

  testConversionOfInvalidMoneyThrowsError() {
    const invalidMoney = {
      amount: 100,
      currency: false,
    }
    const invalidCurrency = '0';
    assert.throws(() => {
      this.bank.convert(invalidMoney, invalidCurrency)
    }, new Error(false + "->" + invalidCurrency))
  }

  getAllTestMethods() {
    let moneyPrototype = BankTest.prototype;
    let allProps = Object.getOwnPropertyNames(moneyPrototype);
    return allProps.filter(p => {
      return typeof moneyPrototype[p] === 'function' && p.startsWith("test");
    });
  }

  runAllTests() {
    let testMethods = this.getAllTestMethods();
    testMethods.forEach(m => {
      console.log("Running: %s()", m);
      var method = Reflect.get(this, m);
      try {
        Reflect.apply(method, this, []);
      } catch (e) {
        if (e instanceof assert.AssertionError) {
          console.log(e);
        } else {
          throw e;
        }
      }
    });
  }
}

new BankTest().runAllTests();
