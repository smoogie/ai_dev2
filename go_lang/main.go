package main

import (
	"ai_dev/c01/lo1"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}
	app := &cli.App{
		Name: "ai_dev",
		Commands: []*cli.Command{
			{
				Name:   "c01l01",
				Action: lo1.C01L01,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
