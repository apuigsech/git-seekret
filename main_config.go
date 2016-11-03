package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
	"strings"
)

func GitSeekretConfig(c *cli.Context) error {
	if c.Bool("init") {
		err := gs.InitConfig()
		if err != nil {
			return err
		}
	} else {
		err := gs.LoadConfig(true)
		if git.IsErrorClass(err, git.ErrClassConfig) {
			return fmt.Errorf("Config not initialised - Try: 'git-seekret config --init'")
		}
		if err != nil {
			return err
		}
	}

	set := c.String("set")

	if set != "" {
		a := strings.SplitN(set, "=", 2)
		key := a[0]
		value := ""
		if len(a) == 2 {
			value = a[1]
			fmt.Println("Value:", value)
		}

		err := setConfig(gs.config, key, value)
		if err != nil {
			return err
		}
	}

	gs.SaveConfig()

	showConfig(gs.config)

	return nil
}

func setConfig(config *gitSeekretConfig, key string, value interface{}) error {
	switch key {
	case "version":
		return fmt.Errorf("not suported")
	case "rulespath":
		rulespath, ok := value.(string)
		if !ok {
			return fmt.Errorf("invalid format")
		}
		config.rulespath = rulespath
	case "rulesenabled":
		return fmt.Errorf("not suported - change enabled rules using 'git-seekret rules'")
	case "exceptionsfile":
		exceptionsfile, ok := value.(string)
		if !ok {
			return fmt.Errorf("invalid format")
		}
		config.exceptionsfile = exceptionsfile
	}
	return nil
}

func showConfig(config *gitSeekretConfig) {
	fmt.Printf("Config:\n")
	fmt.Printf("\tversion = %d\n", config.version)
	fmt.Printf("\trulespath = %s\n", config.rulespath)
	fmt.Printf("\trulesenabled = %s\n", config.rulesenabled)
	fmt.Printf("\texceptionsfile = %s\n", config.exceptionsfile)
}
