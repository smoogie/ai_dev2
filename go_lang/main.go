package main

import (
	"ai_dev/c01/l01/helloapi"
	"ai_dev/c01/l04/blogger"
	"ai_dev/c01/l04/moderation"
	"ai_dev/c01/l05/liar"
	"ai_dev/c02/l02/inprompt"
	"ai_dev/c02/l03/embedding"
	"ai_dev/c02/l04/whisper"
	"ai_dev/c02/l05/functions"
	"ai_dev/c03/l01/rodo"
	"ai_dev/c03/l02/scraper"
	"ai_dev/c03/l03/whoami"
	"ai_dev/c03/l04/search"
	"ai_dev/c03/l05/people"
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
			{
				Name:   "embedding",
				Action: embedding.Command,
			},
			{
				Name:   "whisper",
				Action: whisper.Command,
			},
			{
				Name:   "functions",
				Action: functions.Command,
			},
			{
				Name:   "rodo",
				Action: rodo.Command,
			},
			{
				Name:   "scraper",
				Action: scraper.Command,
			},
			{
				Name:   "whoami",
				Action: whoami.Command,
			},
			{
				Name:   "search",
				Action: search.Command,
			},
			{
				Name:   "people",
				Action: people.Command,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
