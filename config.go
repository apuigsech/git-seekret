package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"strings"
	"github.com/apuigsech/seekret"
	"github.com/apuigsech/seekret/models"
)

type gitSeekretConfig struct {
	configLevel git.ConfigLevel
	gitConfig *git.Config

	version int32
	rulespath string
	rulesenabled string
	exceptionsfile string
}


func (gs *gitSeekret)InitConfig(configLevel git.ConfigLevel) (error) {
	gitConfig, err := openGitConfig(configLevel, gs.repo)
	if err != nil {
		return err
	}
	defer gitConfig.Free()

	gs.config = &gitSeekretConfig{
		configLevel: configLevel,
		gitConfig: gitConfig,
		version: 1,
		rulespath: seekret.DefaultRulesPath(),
		rulesenabled: "",
		exceptionsfile: "",
	}

	err = gs.RunConfig()
	if err != nil {
		return err
	}

	return nil
}


func (gs *gitSeekret)LoadConfig(configLevel git.ConfigLevel, run bool) (error) {
	gitConfig, err := openGitConfig(configLevel, gs.repo)
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
		configLevel: configLevel,
		gitConfig: gitConfig,
		version: version,
		rulespath: rulespath,
		rulesenabled: rulesenabled,
		exceptionsfile: exceptionsfile,
	}

	if run {
		err = gs.RunConfig()
		if err != nil {
			return err
		}
	}

	return nil
}


func (gs *gitSeekret)SaveConfig() (error) {
	if gs.config.gitConfig == nil {
		return fmt.Errorf("git config not loaded")
	}

	err := gs.config.gitConfig.SetInt32("gitseekret.version", gs.config.version)
	if err != nil {
		return err
	}

	err = gs.config.gitConfig.SetString("gitseekret.rulespath", gs.config.rulespath)
	if err != nil {
		return err
	}

	err = gs.config.gitConfig.SetString("gitseekret.rulesenabled", buildRulesEnabledString(gs.seekret.ListRules()))
	if err != nil {
		return err
	}

	err = gs.config.gitConfig.SetString("gitseekret.exceptionsfile", gs.config.exceptionsfile)
	if err != nil {
		return err	
	}

	return nil
}

func (gs *gitSeekret)RunConfig() (error) {
	// TODO: Relative path from repo root.
	err := gs.seekret.LoadRulesFromPath(gs.config.rulespath, false)
	if err != nil {
		return err
	}

	for _,rule := range strings.Split(gs.config.rulesenabled, ",") {
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
		gitConfig, err = git.OpenOndisk(nil, configFile)
		if err != nil {
			return nil, err
		}
	}

	return gitConfig, nil
}


func buildRulesEnabledString(listRules []models.Rule) string {
	rulesenabled := make([]string, 0, len(listRules))
	for _,rule := range listRules {
		if rule.Enabled {
			rulesenabled = append(rulesenabled, rule.Name)
		}
	}
	return strings.Join(rulesenabled, ",")
}
