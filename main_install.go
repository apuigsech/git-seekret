package main

import (
	"github.com/urfave/cli"
)

func GitSeekretInstall(c *cli.Context) error {
	gsc := gitSeekretDefaultConfig
	gsc.gitConfig = gitSeekretCurrentConfig.gitConfig
	gsc.SaveConfig()
	return nil
}