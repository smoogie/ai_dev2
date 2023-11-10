package main

import (
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
				Name:   "fill-data",
				Action: Command,
			},
			{
				Name:   "create-qdrant-collection",
				Action: CreateQdrantCollection,
			},
			{
				Name:   "create-mysql-table",
				Action: CreateMySQLTable,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
