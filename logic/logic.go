// Package logic provides logic-gates-related algorithms.
package logic

// HalfAdder adds two integers using half adder logic.
func HalfAdder(a, b uint8) (carry, sum uint8) {
	return a & b, a ^ b
}

// FullAdder adds two integers using full adder logic.
func FullAdder(a, b, carryIn uint8) (carryOut, sum uint8) {
	return a&b | carryIn&(a^b), a ^ b ^ carryIn
}

// FullSubtractor subtracts two integers (a-b) using full subtractor logic.
func FullSubtractor(a, b, borrowIn uint8) (borrowOut, difference uint8) {
	return ^a&b | borrowIn&^(a^b), a ^ b ^ borrowIn
}
