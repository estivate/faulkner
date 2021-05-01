# Faulkner

A small and easy-to-use logging module for Go programs. It's focused on:

* Simplicity ~ easily for devs to use w/o taking up space in brains
* Consistent ~ various app logs can all look the same for sys admins

It's basically a thin wrapper over Go's `log` package that enables a few
additional things that I find useful:

* Three types of messages: ERROR, INFO and DEBUG
* Easily disable DEBUG (and INFO) messages for live deployments
* ERROR highlighted in red in console (automatically disabled for Windows)
* All messages include date and time
* ERROR messages also include file and line number
* Print simple "banner" messages (for instance, special info at startup)

It logs to Stderr by default, but you can configure it to log to a file
or any other io.Writer. Personally I find it useful to just log to Stderr
and let folks deploying the app decide where to pipe the logs. This way 
we don't need a new binary when those locations change.

If you need fully featured, structured logging in Go check out 
[Logrus](https://github.com/Sirupsen/logrus) or the tons of options
on the logging section of
[Awesome Go](https://github.com/avelino/awesome-go#logging).

If you use Faulkner, you should likely plan to eventually fork it for
your own team. Logging is core to every app you build, and there's 
only a few good reasons to pull in 3rd party code to do it, especially
if you are working in highly regulated industries.

## STATUS

Early version, stay tuned. The API is still changing so if you use it
lock to a version. Only tested on Linux.

## Installation

```go
go get github.com/estivate/faulkner@latest
```

## QUICK START

```go
package main

import (
	"fmt"
	"github.com/estivate/faulkner"
)

func main() {

	// let's make this interesting with some fake data
	app_version := "1.10.4"
	debug_mode_enabled := true
	var_thing := "user did this, for reals"

	// create a default logger
	logger := faulkner.DefaultLogger()

	// format a startup message
	message := fmt.Sprintf("Starting MyApp, version %s.", app_version)
	if debug_mode_enabled {
		message += "\nRunning in Debug Mode"
	}
	logger.PrintBanner(message)
	logger.Debug.Printf("Can you believe a debugging %s.", var_thing)
	logger.Info.Println("Just a normal thing is happening.")
	logger.Error.Printf("Alert! An unexpected %s.", var_thing)

	// turn off DEBUG messages
	logger.DebugOff()
	logger.Debug.Println("This line won't print now.")

	// turn off INFO messages
	logger.InfoOff()
	logger.Debug.Println("This line still won't print.")
	logger.Info.Println("This line won't print now.")
	logger.Error.Println("This is all that shows up!")

}
```
If you run the above your output will look something like this:

```
>> go run main.go
--------------------------
Starting MyApp, version 1.10.4.
Running in Debug Mode
--------------------------
DEBUG: 2021/05/01 08:35:43 example.go:25: Can you believe a debugging user did this, for reals.
INFO:  2021/05/01 08:35:43 Just a normal thing is happening.
ERROR: 2021/05/01 08:35:43 example.go:27: Alert! A unexpected user did this, for reals.
ERROR: 2021/05/01 08:35:43 example.go:37: This is all that shows up!
```
The words "ERROR" will be red in the output above.

## Advanced Usage

In most cases you should probably just fork this repo and modify the defaults so you can keep things
simple and just use the above code too. However, there are some things I occassionally have to do
(like logging to a file or service), and it's not too hard.

Instead of calling `DefaultLogger()` you can call `NewLogger()` and pass it 
[functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
as needed.

```go

// toggle debug messages
ourDebugState = false
logger, _ := faulkner.NewLogger(faulkner.SetDebug(ourDebugState))
logger.Debug.Printf("This message won't appear in log.")

// log to a file
logger, err := faulkner.NewLogger(faulkner.SetFile("/path/to/logfile.txt"))
if err != nil {
    fmt.Printf("Log file can't be written to: %v" err)
}

// pass in multiple options
debug_toggle := faulkner.SetDebug(false)
info_toggle := faulkner.SetInfo(false)
log_file := faulkner.SetFile("/path/to/logfile.txt")
logger, err := faulkner.NewLogger(debug_toggle, info_toggle, log_file)
if err != nil {
    fmt.Printf("Log file can't be written to: %v" err)
}

```
Note that `NewLogger()` can return an error, and in most cases you'll want to check for it.

In the above examples, `logger.Info`, `logger.Debug` and `logger.Error` are instances of
Go's own `*log.Logger` so you can call any method you would on a regular
Go Logger (see [Go Docs](https://golang.org/pkg/log/)).
