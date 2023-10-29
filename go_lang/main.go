package main

import (
	"ai_dev/c01/l04/blogger"
	"ai_dev/c01/l04/moderation"
	lo5 "ai_dev/c01/l05"
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
			{
				Name:   "c01l04_blogger",
				Action: blogger.C01L04_blogger,
			},
			{
				Name:   "c01l04_moderation",
				Action: moderation.C01L04_moderation,
			},
			{
				Name:   "c01l05",
				Action: lo5.C01L05,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
