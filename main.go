package main

import (
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
	"log"
	"os"
	"path/filepath"
)

var gs *gitSeekret

func main() {
	app := cli.NewApp()

	app.Name = "git-seekret"
	app.Version = "0.0.1"
	app.Usage = "prevent from committing sensitive information into git repository"

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
			Name:   "config",
			Usage:  "manage configuration seetings",
			Action: GitSeekretConfig,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "init",
					Usage: "initialize configuration with default values",
				},
				cli.StringFlag{
					Name:  "set, s",
					Usage: "set value for specific setting",
					Value: "",
				},
			},
		},
		{
			Name:   "rules",
			Usage:  "manage rules",
			Action: GitSeekretRules,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "enable, e",
					Usage: "enable rules",
					Value: "",
				},
				cli.StringFlag{
					Name:  "disable, d",
					Usage: "disable rule",
					Value: "",
				},
			},
		},
		{
			Name:   "check",
			Usage:  "inspect git repository",
			Action: GitSeekretCheck,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "commit, c",
					Usage: "include commited files. Argument is the number of commits to inspect (0 = all)",
					Value: 0,
				},
				cli.BoolFlag{
					Name:  "staged, s",
					Usage: "include staged files",
				},
			},
		},
		{
			Name:   "hook",
			Usage:  "manage git hooks",
			Action: GitSeekretHook,
			Flags: []cli.Flag{
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
	var repoPath string

	if c.Bool("global") {
		configLevel = git.ConfigLevelGlobal
	} else {
		configLevel = git.ConfigLevelLocal
	}

	repoPath, err = filepath.Abs(".")
	if err != nil {
		log.Panic(err)
	}

	gs, err = NewGitSeekret(repoPath, configLevel)
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func gitSeekretAfter(c *cli.Context) error {
	return nil
}
