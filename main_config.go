package main

import (
	"fmt"
	"github.com/urfave/cli"
	"strings"
)

func GitSeekretConfig(c *cli.Context) error {
	set := c.String("set")

	if set != "" {
		a := strings.SplitN(set, "=", 2)

		switch a[0] {
			case "rulespath":
				gitSeekretCurrentConfig.rulespath = a[1]
			case "rulesenabled":
				gitSeekretCurrentConfig.rulesenabled = a[1]
			default:
				err := fmt.Errorf("%s is not a valid attribute", a[0])
				return err
		}
		
		gitSeekretCurrentConfig.SaveConfig()
	}

	gitSeekretCurrentConfig.ShowConfig()

	return nil
}