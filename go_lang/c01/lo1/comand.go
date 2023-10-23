package lo1

import (
	"ai_dev/base_flow"
	"fmt"
	"github.com/urfave/cli/v2"
)

func C01L01(c *cli.Context) error {
	err := base_flow.RunProcess("helloapi", process)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	return err
}
