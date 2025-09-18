package calc

import "errors"

type Result struct {
	Value float64
	Err   error
}

func newResult(value float64, err error) Result {
	return Result{
		Value: value,
		Err:   err,
	}
}

func add(x float64, y float64) Result {
	var val float64 = x + y
	return newResult(val, nil)
}

func subtract(minuend float64, subtrahend float64) Result {
	return newResult(minuend-subtrahend, nil)

}

func divide(dividend float64, divisor float64) Result {
	if divisor == 0.0 {
		return newResult(0, errors.New("cannot divide by 0"))
	}
	return newResult(float64((dividend / divisor)), nil)
}

func multiply(factor1 float64, factor2 float64) Result {
	return newResult(factor1*factor2, nil)
}

func sqr(root float64, square float64) Result {
	var result_value float64 = root
	res := newResult(0, nil)
	for range int(square - 1) {
		res = multiply(result_value, root)
		result_value = res.Value
	}
	return res
}

func FunctionChooser(input string, value1 float64, value2 float64) Result {
	switch input {
	case "+":
		return add(value1, value2)
	case "-":
		return subtract(value1, value2)
	case "*":
		return multiply(value1, value2)
	case "/", "\\":
		return divide(value1, value2)
	case "^":
		return sqr(value1, value2)
	}
	return newResult(0, errors.New("improper command, no command parsed"))
}
