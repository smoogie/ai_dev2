package knowledge

import (
	"ai_dev/base_flow"
	"fmt"
	"github.com/urfave/cli/v2"
)

func Command(c *cli.Context) error {
	err := base_flow.RunProcess("knowledge", process)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	return err
}
