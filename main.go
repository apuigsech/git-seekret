package main

import (
	"log"
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
	"os"
)

var gs *gitSeekret

func main() {
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
			Name:	"config",
			Usage:  "TBD",
			Action:   GitSeekretConfig,
			Flags: 	[]cli.Flag {
				cli.BoolFlag{
					Name: "init",
					Usage: "TBD",
				},
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
			Action: GitSeekretRules,
			Flags: 	[]cli.Flag {
				cli.StringFlag{
					Name:  "enable, e",
					Usage: "TBD",
					Value: "",
				},
				cli.StringFlag{
					Name:  "disable, d",
					Usage: "TBD",
					Value: "",
				},
			},	
		},
		{
			Name:     "check",
			Usage:    "TBD",
			Action:   GitSeekretCheck,
			Flags: 	[]cli.Flag {
				cli.IntFlag{
					Name:  "commit, c",
					Usage: "inspect commited files. Argument is the number of commits to inspect (0 = all)",
					Value: 0,
				},
				cli.BoolFlag{
					Name:  "staged, s",
					Usage: "inspect staged files",
				},
			},
		},
		{
			Name: "hook",
			Usage: "TBD",
			Action:   GitSeekretHook,
			Flags: 	[]cli.Flag {
				cli.StringFlag{
					Name:  "run, r",
					Usage: "TBD",
					Value: "",
				},				
				cli.StringFlag{
					Name:  "enable, e",
					Usage: "TBD",
					Value: "",
				},
				cli.StringFlag{
					Name:  "disable, d",
					Usage: "TBD",
					Value: "",
				},
				cli.BoolFlag{
					Name:  "enable-all",
					Usage: "TBD",
				},
				cli.BoolFlag{
					Name:  "disable-all",
					Usage: "TBD",
				},
			},	
		},
	}

	app.Before = gitSeekretBefore
	app.After = gitSeekretAfter

	app.Run(os.Args)
}


func gitSeekretBefore(c *cli.Context) error {
	var configLevel git.ConfigLevel
	var err error

	if c.Bool("global") {
		configLevel = git.ConfigLevelGlobal
	} else {
		configLevel = git.ConfigLevelLocal
	}

	_ = configLevel

	gs,err = NewGitSeekret(".")
	if err != nil {
		log.Panic(err)
	}

	return nil
}


func gitSeekretAfter(c *cli.Context) error {
	return nil
}