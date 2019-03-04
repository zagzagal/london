// main.go
// Paul Schuster
// 030119

package main

import (
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

// Ya config globals... IDK
var pool = "4a4b"
var eval = "a&b"
var numberOfTries = 10000
var deckSize = 60
var handSize = 7
var maxMul = 5
var debug = false
var quiet = false

func registerCallbacks() {
	a := js.FuncOf(DeckCheck)
	v := js.FuncOf(Validate)
	js.Global().Set("deckCheck", a)
	js.Global().Set("validateInput", v)
}

func DeckCheck(this js.Value, i []js.Value) interface{} {
	pool := i[0].String()
	eval := i[1].String()
	m := i[2].String()
	maxMul, _ = strconv.Atoi(m)

	res := deckCheck(pool, eval)
	return res
}

func Validate(this js.Value, i []js.Value) interface{} {
	pool := i[0].String()
	eval := i[1].String()
	m := i[2].String()
	_, err := strconv.Atoi(m)
	if err != nil {
		return false
	}

	return inputVal(pool, eval)
}

func main() {
	c := make(chan struct{}, 0)

	println("london mul WASM edition Initialized")
	// register functions
	rand.Seed(time.Now().UTC().UnixNano())
	registerCallbacks()
	<-c
}
