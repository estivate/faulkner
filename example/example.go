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
