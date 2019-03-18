//go:generate go run ./generators/versioning/main.go
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	bootstrap(app)
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func bootstrap(app *cli.App) {
	app.Name = "gosemver"
	app.Usage = "go forth and semver"
	app.Commands = commands(commandBump, commandGet, commandSet, commandVersion)
	app.Version = fmt.Sprintf("%s-%s", Version, Commit)
	app.Action = handleDefault
}
