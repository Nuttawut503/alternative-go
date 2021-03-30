package behavioral_patterns

import "fmt"

/* Define a family of algorithms,
encapsulate each one, and make them interchangeable.
Strategy lets the algorithm vary independently from clients that use it
*/

type Operator interface {
	do(int, int) int
}

type Plus struct{}

func (Plus) do(a, b int) int {
	return a + b
}

type Minus struct{}

func (Minus) do(a, b int) int {
	return a - b
}

type Multiply struct{}

func (Multiply) do(a, b int) int {
	return a * b
}

type Calculator struct {
	currentValue int
	operator     Operator
}

func (cal *Calculator) SetOperator(op Operator) {
	cal.operator = op
}

func (cal *Calculator) ApplyWith(val int) {
	cal.currentValue = cal.operator.do(cal.currentValue, val)
}

func (cal Calculator) GetCurrentValue() int {
	return cal.currentValue
}

func StrategyExample() {
	cal := Calculator{}
	cal.SetOperator(Plus{})
	cal.ApplyWith(5)
	cal.ApplyWith(9)
	cal.SetOperator(Multiply{})
	cal.ApplyWith(2)
	cal.SetOperator(Minus{})
	cal.ApplyWith(9)
	fmt.Println(cal.GetCurrentValue())
}
