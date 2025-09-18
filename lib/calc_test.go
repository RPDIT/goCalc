package calc

import (
	"fmt"
	"testing"
)

type Test struct {
	value1          float64
	value2          float64
	result          Result
	confirmed_value float64
}

func NewTest(value1 float64, value2 float64, testFunc func(float64, float64) Result, result_value float64) Test {
	return Test{
		value1:          value1,
		value2:          value2,
		result:          testFunc(value1, value2),
		confirmed_value: result_value,
	}
}

func CheckValue(phrase1 bool, err error, t *testing.T) {
	if phrase1 {
		t.Errorf("Error: %v", err)
	}
}

func TestAdd(t *testing.T) {
	test_value := NewTest(69, 777, add, 846)
	addition_error := fmt.Errorf(
		"The calculator failed to find %v as the sum of %v + %v. \nReturned: %v",
		test_value.confirmed_value,
		test_value.value1,
		test_value.value2,
		test_value.result.Value)
	CheckValue(
		(test_value.confirmed_value != test_value.result.Value),
		addition_error,
		t)
}

func TestSubtract(t *testing.T) {
	test_value := NewTest(45, 22.5, subtract, 22.5)
	CheckValue(
		test_value.result.Value != test_value.confirmed_value,
		fmt.Errorf(
			"Failed to subtract %v from %v: %v",
			test_value.value1,
			test_value.value2,
			test_value.result.Err),
		t,
	)
}

func TestMultiply(t *testing.T) {
	test_value := NewTest(22.5, -2, multiply, -45)
	CheckValue(
		test_value.result.Value != test_value.confirmed_value,
		fmt.Errorf("Failed to multiply %v and %v.  Received: %v", test_value.value1, test_value.value2, test_value.result.Value),
		t,
	)
}

func TestDivide(t *testing.T) {
	test_value := NewTest(50, 2, divide, 25)
	CheckValue(
		test_value.result.Err != nil,
		fmt.Errorf(
			"Failed to divide %v by %v: %v",
			test_value.value1,
			test_value.value2,
			test_value.result.Err,
		),
		t,
	)
	CheckValue(
		test_value.result.Value != test_value.confirmed_value,
		fmt.Errorf(
			"Failed to divide %v by %v. Got %v",
			test_value.value1,
			test_value.value2,
			test_value.result.Value),
		t,
	)
}

func TestSqr(t *testing.T) {
	test_value := NewTest(10, 3, sqr, 1000)
	CheckValue(
		test_value.result.Value != test_value.confirmed_value,
		fmt.Errorf("Failed to find %v to the %v power.  Received: %v", test_value.value1, test_value.value2, test_value.result.Value),
		t,
	)
}
