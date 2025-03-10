package main

import (
	"test-plus/cmd"
	"test-plus/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupMySQL()

	// run http
	cmd.ServeHTTP()
}
