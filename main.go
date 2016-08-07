package main

import (
	seekret "github.com/apuigsech/seekret/lib"
	"github.com/urfave/cli"
	"os"
)

const (
	GitSeekretBaseFolder = ".git/seekret/"
	GitSeekretEnabledRulesFile = GitSeekretBaseFolder + "xxxx"
)

var s *seekret.Seekret

func main() {
	s = seekret.NewSeekret()

	app := cli.NewApp()

	app.Name = "git-seekret"
	app.Version = "0.0.1"
	app.Usage = "TBD"

	app.Author = "Albert Puigsech Galicia"
	app.Email = "albert@puigsech.com"

	app.EnableBashCompletion = false

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "global",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:     "install",
			Usage:    "TBD",
			Action:   GitSeekretInstall,
		},
		{
			Name:	"config",
			Usage:  "TBD",
			Action:   GitSeekretConfig,
			Flags: 	[]cli.Flag {
				cli.StringFlag{
					Name:  "set, s",
					Usage: "TBD",
					Value: "",
				},
			},
		},
		{
			Name: "rules",
			Usage: "Manage rules",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Usage: "TBD",
					Action: GitSeekretRulesList,
				},
				{
					Name: "enable",
					Usage: "TBD",
					Action: GitSeekretRulesEnable,
				},
				{
					Name: "disable",
					Usage: "TBD",
					Action: GitSeekretRulesDisable,
				},
			},
		},
		{
			Name:     "check",
			Usage:    "TBD",
			Action:   GitSeekretCheck,			
		},
	}

	app.Before = gitSeekretBefore
	app.After = gitSeekretAfter

	app.Run(os.Args)
}


func gitSeekretBefore(c *cli.Context) error {
	gitSeekretCurrentConfig = NewGitSeekretConfig()

	err := gitSeekretCurrentConfig.LoadConfig(c.Bool("global"))
	if err != nil {
		return err
	}

	return nil
}


func gitSeekretAfter(c *cli.Context) error {
	return saveRules()
}