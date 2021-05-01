# Faulkner

A simple and flexible logging solution for Go programs. It's focused on:

* Simplicity ~ new team members can easily use it w/o reading anything
* Consistent ~ all your apps can log the same way

## Features

Faulkner is a thin wrapper over Go's `log` package. It enables a few cool things
in a pretty simple way:

* INFO & Error messages always appear in logs
* DEBUG messages can be toggled on/off
* Errors highlighted with Red typed prefix
* Errors include file and line number
* Outputs to Stderr, a log file, or any other io.Writer
* Print "banners" to find start, stop or other special events in logs

## STATUS

Early version, stay tuned. Only tested on Linux.

## QUICK START

The simpliest way to use is to embrace the defaults. By default all logging will print to Stderr
and all messages (including DEBUG) will print.

```go
package main

import (
    "fmt"
    "github.com/estivate/faulkner"
)

func main() {
    app_version := "1.10.4"
    port_number := ":80"
    logger, _ := faulkner.NewLogger()

    message := fmt.Sprintf("Starting MyApp, version %s.", app_version)
    logger.PrintBanner(message)
    logger.Debug.Printf("Debugging about port %s.", port_number)
    logger.Info.Printf("Info message about port %s.", port_number)
    logger.Error.Printf("Error message thrown regarding port %s.", port_number)
}
```

If you run the above your output will look something like this:

```bash
>> go run main.go
--------------------------
Starting MyApp, version 1.10.4
--------------------------
DEBUG: 2021/04/30 17:15:22 Debugging about port :80.
INFO: 2021/04/30 17:15:22 Info message about port :80.
ERROR: 2021/04/30 17:15:22 main.go:12: Error message thrown regarding port :80.
```
The word "ERROR" will be red in the output above.

If you'd like to toggle DEBUG messages on/off you can pass a bool value to the `SetDebug()` option:

```go
    ourDebugState = false
    logger, _ := faulkner.NewLogger(faulkner.SetDebug(ourDebugState))
	logger.Debug.Printf("This message won't appear in log.")
```

If you'd like to write to a log file instead of Stderr you can provide a file path to the `SetFile()` option:

```go
    logger, err := faulkner.NewLogger(faulkner.SetFile("/path/to/logfile.txt"))
    if err != nil {
        fmt.Printf("Log file can't be written to: %v" err)
    }
```

Note that we are checking the error value here. Default usage or toggling Debug will never error,
but specifying a log file can error if permissions are wrong, etc.

Finally, you can set multiple options:

```go
    debug_toggle := faulkner.SetDebug(false)
    log_file := faulkner.SetFile("/path/to/logfile.txt")
    logger, err := faulkner.NewLogger(debug_toggle, log_file)
    if err != nil {
        fmt.Printf("Log file can't be written to: %v" err)
    }
```

