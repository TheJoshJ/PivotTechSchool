package calculator

type ErrDivByZero struct{}

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) (int, error) {
	divByZero := ErrDivByZero{}
	if b == 0 {
		return 0, divByZero
	}
	return a / b, nil
}

func (ErrDivByZero) Error() string {
	return "divide by zero"
}
