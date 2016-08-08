package main

import (
	"fmt"
	"github.com/urfave/cli"
	seekret "github.com/apuigsech/seekret/lib"
)

func GitSeekretCheck(c *cli.Context) error {
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

	err := s.LoadObjects(seekret.SourceTypeGit, ".", options)
	if err != nil {
		return err
	}

	s.Inspect(4)

	for _,x := range s.ListSecrets() {
		fmt.Printf("%#v\n", x)
	}

	return nil
}