# Digital Clock

This is a very simple example showing how to implement a timer in Go/WebAssembly. I decided to write it up because at this time (November 2018), Go's WebAssembly support is still very new and there are few good tutorial examples online.

### Example in JavaScript

In JavaScript, a simple digital clock can be implemented like this:

```
function UpdateClock()
{
        var now = new Date();
        var elt = document.getElementById("clock");
        elt.innerHTML = now.toLocaleTimeString();
        setTimeout(UpdateClock,1000)
}

window.onload = UpdateClock;
```

The last line calls the `UpdateClock()` function when the page is loaded. The last line in `UpdateClock()` calls `setTimeout()` to call `UpdateClock()` again after one second (1000 milliseconds), so the clock will continue to update.

Within `UpdateClock()`, the time is retrieved by creating a new `Date` object, adjusting the time to the local time zone with `toLocaleTimeString()`, then inserting that into the browser's DOM. In the HTML, there is a tag that looks something like

```
<div id="clock"></div>
```

which has no content in this example. `getElementById()` gets a handle on that tag so it can be modified. Every time the text is updated, the browser updates the display.

### Let's Do That in Go

So how do we do all of that in Go? Let's take it step by step.

Creating a new `Date` object is done like this:

```
now := js.Global().Get("Date").New()
```

`js.Global()` gets the JavaScript global object, and `Get("Date")` returns the `Date` object. `New()` creates a new instance of it. To get the local time,

```
s := now.Call("toLocaleTimeString").String()
```

calls the `toLocaleTimeString()` method of `now`, then converts it into a Go string type. This is what we want to use for the text content in the web page. To set that,

```
js.Global().Get("document").Call("getElementById", "clock").Set("textContent", s)
```

again gets the JavaScript global object, gets its document object, and calls `document.getElementById()` using "`clock`" as the id. Using the returned object, `js.Set()` sets the `textContent` (the same as `innerHTML`) property to the current time.

The last part is to set up a timer event. Rather than using setTimer() every time we update the time, let's set up an interval timer to deliver timer events every 200 ms, and call the update code from that.

```
timer_cb := js.NewCallback(update_time)
```

Creates a callback to `function update_time()` that contains the above Go code. Then

```
js.Global().Call("setInterval",timer_cb,"200")
```

calls the JavaScript `setInterval()` function to run the callback every 200 milliseconds.

#### Putting it Together

Here is the full Go code:

```
func update_time(args []js.Value) {
        s := js.Global().Get("Date").New().Call("toLocaleTimeString").String()
        js.Global().Get("document").Call("getElementById", "clock").Set("textContent", s)
}
```

and in main(),

```
timer_cb := js.NewCallback(update_time)
js.Global().Call("setInterval",timer_cb,"200")
```

As you can see, writing "JavaScript in Go" is more verbose and complicated because we need to use Go's `js` package to access JavaScript from Go. But functions and methods in `js` provide ways for us to do everything in Go that can be done in JavaScript, and there is a nearly 1:1 correspondence between the JavaScript code and what we need to write in Go. So although it's a little more complicated, it's really not difficult at all.

### Using the repository

You can see the actual code (with some minor differences compared to this README file) in `main.go`.

The following files are necessary for deployment:

##### For basic functionality

```
index.html	- the web page
main.wasm	- the compiled Go code
wasm_exec.js	- standard JavaScript glue code
```

##### For appearance (CSS and images)

```
backgnd_tile.gif
favicon.ico
styles.css
```

#### Deployment

Just copy all of the files in the above two groups to a directory on your web server.
To see the clock run, load index.html into your web browser.

A simple web server written in Go is included. To view the page
on your computer, start the web server, like this

```
$ go run webserver.go
2018/11/20 12:43:02 listening on ":8080"...
```

then direct your browser to http://localhost:8080

### Author

	Jay Ts
	(http://jayts.com)

### Copyright

	Copyright 2018 Jay Ts

	Released under the GNU Public License, version 3.0 (GPLv3)
	(http://www.gnu.org/licenses/gpl.html)
