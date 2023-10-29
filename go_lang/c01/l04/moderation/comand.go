package moderation

import (
	"ai_dev/base_flow"
	"fmt"
	"github.com/urfave/cli/v2"
)

func C01L04_moderation(c *cli.Context) error {
	err := base_flow.RunProcess("moderation", process)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	return err
}
