package app

import "github.com/urfave/cli"

// Generate will return the application cli to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search for IPs and Server names on internet"

	return app
}