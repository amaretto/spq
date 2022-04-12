package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func executeQuery(c *cli.Context) error {
	var args = c.Args().Slice()
	if len(args) == 0 {
		return fmt.Errorf("no target args specified. see `spq query -h` for more details")
	} else if len(args) > 1 {
		return fmt.Errorf("too many args specified. see `spq query -h` for more details")
	}
	return nil
}

func query() {
	// ToDo: implement
	fmt.Println("exec query")
}
