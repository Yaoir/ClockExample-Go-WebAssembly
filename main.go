// Digital Clock - Shows use of setInterval() in Go/WebAssembly
// build: GOOS=js GOARCH=wasm go build -o main.wasm main.go

package main

import "syscall/js"

// Callback for the interval timer
// Get the current time and update it in the DOM

func update_time(this js.Value, args []js.Value) interface{} {
	// Get the current date in this locale
	// It's done like this in JavaScript:
	//	date = new Date()
	//	s = date.toLocaleTimeString()

	date := js.Global().Get("Date").New()
	s := date.Call("toLocaleTimeString").String()

	// update the text in <div id="clock">
	// It's done like this in JavaScript:
	//	document.getElementById("clock").textContent = s

	js.Global().Get("document").Call("getElementById", "clock").Set("textContent", s)
	return nil
}

func main() {
	// Set up a recurring timer event to call update_time() every 200 ms.
	// It's done like this in JavaScript:
	//	setInterval(update_time,200)

	// Create JavaScript callback connected to update_time()

	timer_cb := js.FuncOf(update_time)

	// Set timer to call timer_cb() every 200 ms.

	js.Global().Call("setInterval",timer_cb,"200")

	// An empty select blocks, so the main() function will never exit.
	// This allows the event handler callbacks to continue operating.
	select{}
}
