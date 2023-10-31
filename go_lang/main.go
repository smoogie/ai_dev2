package main

import (
	"ai_dev/c01/l01/helloapi"
	"ai_dev/c01/l04/blogger"
	"ai_dev/c01/l04/moderation"
	"ai_dev/c01/l05/liar"
	"ai_dev/c02/l02/inprompt"
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
				Name:   "helloapi",
				Action: helloapi.Command,
			},
			{
				Name:   "blogger",
				Action: blogger.Command,
			},
			{
				Name:   "moderation",
				Action: moderation.Command,
			},
			{
				Name:   "liar",
				Action: liar.Command,
			},
			{
				Name:   "inprompt",
				Action: inprompt.Command,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
