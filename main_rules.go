package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
)

func GitSeekretRules(c *cli.Context) error {
	err := gs.LoadConfig(true)
	if git.IsErrorClass(err, git.ErrClassConfig) {
		return fmt.Errorf("Config not initialised - Try: 'git-seekret config --init'")
	}
	if err != nil {
		return err
	}

	enable := c.String("enable")
	disable := c.String("disable")

	enableAll := c.Bool("enable-all")
	disableAll := c.Bool("disable-all")

	if enableAll {
		fmt.Println("Enabling all rules.")
		gs.EnableRule(".*")
	} else if disableAll {
		fmt.Println("Disabling all rules.")
		gs.DisableRule(".*")
	}

	// "all" represents that either enableAll or disableAll was triggered.
	all := (enableAll || disableAll)

	// If neither enableAll nor disableAll were used, let's look into doing
	// individual enable and/or disable operations.
	// Useful because in a single command you can specify both an --enable flag
	// and a --disable flag but only want to do it if neither enable-all or disable-all were used.
	if !all && enable != "" {
		gs.EnableRule(enable)
	}

	if !all && disable != "" {
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
