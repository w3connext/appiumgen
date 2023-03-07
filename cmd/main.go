package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/w3connext/appiumgen/internal/appiumgen"
	"github.com/w3connext/appiumgen/pkg/config"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "platform",
				Aliases: []string{"p"},
				Value:   "",
				Usage:   "Platform name [flutter, android, ios], Please select one.",
			},
		},
		Action: func(c *cli.Context) error {
			platform := c.String("platform")
			if platform == "flutter" || platform == "android" || platform == "ios" {
				appiumgen.Run(config.Config{
					Platform: platform,
				})
				return nil
			}
			return errors.New(fmt.Sprintf("Unsupported platform %s", platform))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
