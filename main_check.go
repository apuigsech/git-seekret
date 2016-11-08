package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
)

func GitSeekretCheck(c *cli.Context) error {
	err := gs.LoadConfig(true)
	if git.IsErrorClass(err, git.ErrClassConfig) {
		return fmt.Errorf("Config not initialised - Try: 'git-seekret config --init'")
	}
	if err != nil {
		return err
	}

	options := map[string]interface{}{
		"commit-files": false,
		"staged-files": false,
	}

	if c.IsSet("commit") {
		options["commit-files"] = true
		options["commit-messages"] = true
		options["commit-count"] = c.Int("commit")
	}

	if c.IsSet("staged") {
		options["staged-files"] = true
	}

	secrets, err := gs.RunCheck(options)
	if err != nil {
		return err
	}
	if secrets != 0 {
		return fmt.Errorf("Please remove discovered secrets")
	}

	return nil
}
