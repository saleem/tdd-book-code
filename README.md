This is a [monorepo](https://trunkbaseddevelopment.com/monorepos/) containing all the source code for the ["Learning Test-Driven Development"](https://learning.oreilly.com/library/view/learning-test-driven-development/9781098106461/) book, published by O'Reilly.

# Prerequisites

You need to install the runtime environments for [Go](https://golang.org/), [Node.js](https://nodejs.org/en/), and [Python 3](https://www.python.org/) to run the code in this repo.

Beyond these, you may need other tools -- e.g. ["act"](https://github.com/nektos/act) to run the GitHub actions locally, [gocyclo](https://github.com/fzipp/gocyclo) to check cyclomatic complexity of Go, [jshint](https://jshint.com) to do static analysis of JavaScript code, and [flake8](https://flake8.pycqa.org) for static analysis of Python code.

To understand the evolution and purpose of the code, you need the accompanying book.

# How to run tests
In brief, use the following commands to run the tests for each language.

## Go
```
cd go
go test -v ./...
```

## JavaScript
```
node js/test_money.js
```

## Python
```
python3 py/test_money.py
```
# How to get the book

The book is available at [Amazon](https://www.amazon.com/Learning-Test-Driven-Development-Polyglot-Uncluttered/dp/1098106474/ref=sr_1_3?) and [O'Reilly](https://learning.oreilly.com/library/view/learning-test-driven-development/9781098106461/).