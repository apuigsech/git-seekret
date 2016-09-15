package main

import (
	"fmt"
	"github.com/apuigsech/seekret-source-git"
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
	"runtime"
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
		"commit": false,
		"staged": false,
	}

	if c.IsSet("commit") {
		options["commit"] = true
		options["commitcount"] = c.Int("commit")
	}

	if c.IsSet("staged") {
		options["staged"] = true
	}

	err = gs.seekret.LoadObjects(sourcegit.SourceTypeGit, gs.repo, options)
	if err != nil {
		return err
	}

	gs.seekret.Inspect(runtime.NumCPU())

	listSecrets := gs.seekret.ListSecrets()
	fmt.Printf("Found Secrets: %d\n", len(listSecrets))
	for _, s := range listSecrets {
		fmt.Printf("\t%s:%d\n", s.Object.Name, s.Nline)
		fmt.Printf("\t\t- Metadata:\n")
		for k, v := range s.Object.Metadata {
			fmt.Printf("\t\t  %s: %s\n", k, v)
		}
		fmt.Printf("\t\t- Rule:\n\t\t  %s\n", s.Rule.Name)

		fmt.Printf("\t\t- Content:\n\t\t  %s\n", s.Line)
	}

	return nil
}
