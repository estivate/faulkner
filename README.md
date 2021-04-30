# Faulkner

A simple and flexible logging solution for Go programs.

## Features

Faulkner is a thin wrapper over Go's `log` package. It enables a few cool things
in a pretty simple way:

* log to Stderr or to a file (or any other io.Writer)
* Offers DEBUG, INFO and ERROR log types
* ERROR messages in red when using Stderr
* ERROR messages include file name and line number
* DEBUG messages can be toggled off for production

## QUICK START

The simpliest way to use is to embrace the defaults. By default all logging will print to Stderr
and all messages (including DEBUG) will print.

```go
package main

import (
    "github.com/estivate/faulkner"
)

func main() {
    logger, _ = faulkner.NewLogger()

	logger.LogDebug.Printf("This is a debug log test.")
	logger.LogInfo.Printf("This is an info log test.")
    logger.LogError.Printf("This is an error log test.")
}
```

If you run the above your output will look something like this:

```
DEBUG: 2021/04/30 17:15:22 This is a debug log test.
INFO: 2021/04/30 17:15:22 This is an info log test.
ERROR: 2021/04/30 17:15:22 main.go:12: This is an error log test.
```
The word <span style="color:blue">Error</span> should be red in the output above.

If you'd like to toggle DEBUG messages on/off you can pass a bool value to the `SetDebug()` option:

```go
    ourDebugState = false
    logger, _ = faulkner.NewLogger(faulkner.SetDebug(ourDebugState))
	logger.LogDebug.Printf("This message won't appear in log.")
```

If you'd like to write to a log file instead of Stderr you can provide a file path to the `SetFile()` option:

```go
    logger, err = faulkner.NewLogger(faulkner.SetFile("/path/to/logfile.txt"))
    if err != nil {
        fmt.Printf("Log file can't be written to: %v" err)
    }
```

Note that we are checking the error value here. Default usage or toggling Debug will never error,
but specifying a log file can error if permissions are wrong, etc.

Finally, you can send multipe options:

```go
    debug_toggle := faulkner.SetDebug(false)
    log_file := faulkner.SetFile("/path/to/logfile.txt")
    logger, err = faulkner.NewLogger(debug_toggle, log_file)
    if err != nil {
        fmt.Printf("Log file can't be written to: %v" err)
    }
```