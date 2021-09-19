class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency

    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)

    def divide(self, divisor):
        return Money(self.amount / divisor, self.currency)

    def __eq__(self, other):
        return self.amount == other.amount and self.currency == other.currency

    def __str__(self):
        return f"{self.currency} {self.amount:0.2f}"

    def __add__(self, a):
        if a is not None and self.currency == a.currency:
            return Money(self.amount + a.amount, self.currency)
        else:
            return None
