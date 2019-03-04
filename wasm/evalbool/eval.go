// eval.go
// Paul Schuster
// 030119

package evalbool

import (
	"container/list"
	"errors"
	"fmt"
)

var ErrCharIll = errors.New("Illegal character in eval expression")
var ErrEval = errors.New("Malformed eval expression")

func pop(l *list.List) bool {
	e := l.Back()
	l.Remove(e)
	return e.Value.(bool)
}

func push(l *list.List, b bool) {
	l.PushBack(b)
}

// EvalRPN takes a byte array in reverse polish notation and evaluates its boolian
// value
func EvalRPN(exp []byte) (bool, error) {
	s := list.New()
	malformed := false
	check := func(i int) bool {
		if s.Len() < i {
			malformed = true
			return false
		}
		return true
	}

	for _, v := range exp {
		switch v {
		case '!':
			if check(1) {
				a := pop(s)
				push(s, !a)
			}
		case '&':
			if check(2) {
				a := pop(s)
				b := pop(s)
				push(s, a && b)
			}
		case '|':
			if check(2) {
				a := pop(s)
				b := pop(s)
				push(s, a || b)
			}
		case '^':
			if check(2) {
				a := pop(s)
				b := pop(s)
				push(s, a != b)
			}
		case 'T', 't':
			push(s, true)
		case 'F', 'f':
			push(s, false)
		default:
			return false, ErrCharIll
		}
		if malformed {
			return false, ErrEval
		}
	}
	if s.Len() != 1 {
		fmt.Println(string(exp))
		fmt.Println(s)
		return false, ErrEval
	}
	return pop(s), nil
}

// Eval takes a boolian expression in infix notation and evaluates it boolian value
func Eval(exp []byte) (bool, error) {
	e := Shunt(exp)
	return EvalRPN(e)
}
