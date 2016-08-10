package main

import (
	"github.com/urfave/cli"
)

func GitSeekretInstall(c *cli.Context) error {
	gs.InitConfig()
	gs.SaveConfig()
	return nil
}