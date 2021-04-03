const assert = require('assert');
const Money = require('./money');
const Portfolio = require('./portfolio');

class MoneyTest {
  testMultiplication() {
    let tenEuros = new Money(10, "EUR");
    let twentyEuros = new Money(20, "EUR");
    assert.deepStrictEqual(tenEuros.times(2), twentyEuros);
  }
  testDivision() {
    let originalMoney = new Money(4002, "KRW")
    let expectedMoneyAfterDivision = new Money(1000.5, "KRW")
    assert.deepStrictEqual(originalMoney.divide(4), expectedMoneyAfterDivision)
  }

  testAddition() {
    let fiveDollars = new Money(5, "USD");
    let tenDollars = new Money(10, "USD");
    let fifteenDollars = new Money(15, "USD");
    let portfolio = new Portfolio();
    portfolio.add(fiveDollars, tenDollars);
    assert.deepStrictEqual(portfolio.evaluate("USD"), fifteenDollars);
  }

  testAdditionOfDollarsAndEuros() {
    var fiveDollars = new Money(5, "USD");
    var tenEuros = new Money(10, "EUR");
    var portfolio = new Portfolio();
    portfolio.add(fiveDollars, tenEuros);
    var expectedValue = new Money(17, "USD");
    assert.deepStrictEqual(portfolio.evaluate("USD"), expectedValue);
  }

  getAllTestMethods() {
    let moneyPrototype = MoneyTest.prototype;
    let allProps = Object.getOwnPropertyNames(moneyPrototype);
    let testMethods = allProps.filter(p => {
      return typeof moneyPrototype[p] === 'function' && p.startsWith("test"); // <3>
    });
    return testMethods;
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

new MoneyTest().runAllTests();
