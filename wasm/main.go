// main.go
// Paul Schuster
// 030119

package main

import (
	"fmt"
	"strconv"
	"syscall/js"
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

func output(b bool) {
	js.Global().Get("document").Call("getElementById", "runButton").Set("result", b)
}

func updateProg(i, n int) {
	pb := js.Global().Get("document").Call("getElementById",
		"prog")
	pb.Set("value", i)
	pb.Set("max", n)
}

func GetData() {
	p := js.Global().Get("document").Call("getElementById",
		"pool").Get("value").String()
	e := js.Global().Get("document").Call("getElementById",
		"eval").Get("value").String()
	pool = p
	eval = e
}

func registerCallbacks() {
	s := js.FuncOf(Stuff)
	js.Global().Set("start", js.FuncOf(start))
	js.Global().Set("stuff", s)
}

func Stuff(this js.Value, i []js.Value) interface{} {
	pool := i[0].String()
	eval := i[1].String()
	m := i[2].String()
	maxMul, _ = strconv.Atoi(m)

	res := doStuff(pool, eval)
	return res
}

func main() {
	c := make(chan struct{}, 0)

	println("london mul WASM edition Initialized")
	// register functions
	registerCallbacks()
	<-c
}

func start(this js.Value, i []js.Value) interface{} {
	GetData()
	// print the header
	if !quiet {
		fmt.Printf("London muligan sim\n")
		fmt.Printf("Deck Size = %d\n", deckSize)
		fmt.Printf("Hand Size = %d\n", handSize)
		fmt.Printf("Min after muligan hand size = %d\n\n", maxMul)
		fmt.Printf("Deck String: %s\n", pool)
		fmt.Printf("eval String: %s\n", eval)
	}

	// do the thing that you do
	process(pool, eval)
	return nil
}
