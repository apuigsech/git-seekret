package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
)

func GitSeekretRules(c *cli.Context) error {
	// TODO: Implement also support for --global
	err := gs.LoadConfig(git.ConfigLevelLocal, true)
	if git.IsErrorClass(err, git.ErrClassConfig) {
		return fmt.Errorf("Config not initialised - Try: 'git-seekret config --init'")
	}
	if err != nil {
		return err
	}

	enable := c.String("enable")
	disable := c.String("disable")

	if enable != "" {
		gs.EnableRule(enable)
	}

	if disable != "" {
		gs.DisableRule(disable)
	}

	fmt.Println("List of rules:")
	for _, r := range gs.seekret.ListRules() {
		status := " "
		if r.Enabled {
			status = "x"
		}
		fmt.Printf("\t[%s] %s\n", status, r.Name)
	}

	gs.SaveConfig()

	return nil
}
