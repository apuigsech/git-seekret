package main

import (
	"errors"
	"fmt"
	"github.com/apuigsech/seekret"
	"github.com/apuigsech/seekret/models"
	"github.com/libgit2/git2go"
	"os"
	"path/filepath"
	"strings"
)

type gitSeekretConfig struct {
	version        int32
	rulespath      string
	rulesenabled   string
	exceptionsfile string
}

func (gs *gitSeekret) InitConfig() error {
	gitConfig, err := openGitConfig(gs.configLevel, gs.repo)
	if err != nil {
		return err
	}
	defer gitConfig.Free()

	gs.config = &gitSeekretConfig{
		version:        1,
		rulespath:      seekret.DefaultRulesPath(),
		rulesenabled:   "",
		exceptionsfile: "",
	}

	err = gs.CheckConfig()
	if err != nil {
		return err
	}

	err = gs.RunConfig()
	if err != nil {
		return err
	}

	return nil
}

func (gs *gitSeekret) CheckConfig() error {
	if gs.config == nil {
		return errors.New("Undefined config file.")
	}
	_, err := os.Stat(gs.config.rulespath)
	if err != nil {
		// If we are using the default rules path, let the user know that there is a SEEKRET_RULES_PATH env var.
		// We don't want to tell the user to use the default path and to instead override it with the
		// SEEKRET_RULES_PATH env var.
		if gs.config.rulespath == seekret.DefaultRulesPath() || os.Getenv("SEEKRET_RULES_PATH") == "" {
			return fmt.Errorf("Unable to use default rulespath \"%s\".\nSystem Error: %s\n\nHOW TO FIX:\n"+
				"Create your own rules folder and set the path of the folder to"+
				"\"SEEKRET_RULES_PATH\" in your environment to override the default rules path.\n"+
				"Example command to create folder:\n$ mkdir -p $HOME/.seekret_rules && "+
				"export $SEEKRET_RULES_PATH=$HOME/.seekret_rules\n"+
				"Reinitialize your config afterwards.\n", gs.config.rulespath, err.Error())
		} else {
			// If something goes really wrong, just return a generic error.
			return fmt.Errorf("Unable to use rulespath \"%s\". Ensure it exists and you"+
				"have permission to read from it. Reinitialize your config afterwards.\n"+
				"System Error: %s\n", gs.config.rulespath, err.Error())
		}
	}
	return nil
}

func (gs *gitSeekret) LoadConfig(run bool) error {
	gitConfig, err := openGitConfig(gs.configLevel, gs.repo)
	if err != nil {
		return err
	}
	defer gitConfig.Free()

	version, err := gitConfig.LookupInt32("gitseekret.version")
	if err != nil {
		return err
	}

	rulespath, err := gitConfig.LookupString("gitseekret.rulespath")
	if err != nil {
		return err
	}

	rulesenabled, err := gitConfig.LookupString("gitseekret.rulesenabled")
	if err != nil {
		return err
	}

	exceptionsfile, err := gitConfig.LookupString("gitseekret.exceptionsfile")
	if err != nil {
		return err
	}

	gs.config = &gitSeekretConfig{
		version:        version,
		rulespath:      rulespath,
		rulesenabled:   rulesenabled,
		exceptionsfile: exceptionsfile,
	}

	if run {
		err = gs.CheckConfig()
		if err != nil {
			return err
		}
		err = gs.RunConfig()
		if err != nil {
			return err
		}
	}

	return nil
}

func (gs *gitSeekret) SaveConfig() error {
	gitConfig, err := openGitConfig(gs.configLevel, gs.repo)
	if err != nil {
		return err
	}
	defer gitConfig.Free()

	err = gitConfig.SetInt32("gitseekret.version", gs.config.version)
	if err != nil {
		return err
	}

	err = gitConfig.SetString("gitseekret.rulespath", gs.config.rulespath)
	if err != nil {
		return err
	}

	err = gitConfig.SetString("gitseekret.rulesenabled", buildRulesEnabledString(gs.seekret.ListRules()))
	if err != nil {
		return err
	}

	err = gitConfig.SetString("gitseekret.exceptionsfile", gs.config.exceptionsfile)
	if err != nil {
		return err
	}

	return nil
}

func (gs *gitSeekret) RunConfig() error {
	// TODO: Relative path from repo root.
	err := gs.seekret.LoadRulesFromPath(gs.config.rulespath, false)
	if err != nil {
		return err
	}

	for _, rule := range strings.Split(gs.config.rulesenabled, ",") {
		gs.seekret.EnableRule(rule)
	}

	// TODO: Relative path from repo root.
	if gs.config.exceptionsfile != "" {
		err := gs.seekret.LoadExceptionsFromFile(gs.config.exceptionsfile)
		if err != nil {
			return err
		}
	}

	return nil
}

func openGitConfig(configLevel git.ConfigLevel, repo string) (*git.Config, error) {
	var gitConfig *git.Config
	var err error

	if configLevel == git.ConfigLevelLocal {
		r, err := git.OpenRepositoryExtended(repo, git.RepositoryOpenCrossFs, "")
		if err != nil {
			return nil, err
		}

		gitConfig, err = r.Config()
		if err != nil {
			return nil, err
		}
	} else {
		var configFile string
		switch configLevel {
		case git.ConfigLevelSystem:
			configFile, err = git.ConfigFindSystem()
		case git.ConfigLevelGlobal:
			configFile, err = git.ConfigFindGlobal()
		case git.ConfigLevelXDG:
			configFile, err = git.ConfigFindXDG()
		}
		if err != nil {
			return nil, err
		}
		configFile, err = filepath.EvalSymlinks(configFile)
		if err != nil {
			return nil, err
		}
		configFile, err = filepath.Abs(configFile)
		if err != nil {
			return nil, err
		}
		gitConfig, err = git.OpenOndisk(nil, configFile)
		if err != nil {
			return nil, err
		}
	}

	return gitConfig, nil
}

func buildRulesEnabledString(listRules []models.Rule) string {
	rulesenabled := make([]string, 0, len(listRules))
	for _, rule := range listRules {
		if rule.Enabled {
			rulesenabled = append(rulesenabled, rule.Name)
		}
	}
	return strings.Join(rulesenabled, ",")
}
