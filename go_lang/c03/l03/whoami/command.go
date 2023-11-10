package whoami

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Command(c *cli.Context) error {
	err := process()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	return err
}
