package main

import (
	"fmt"
	"path/filepath"
	"github.com/libgit2/git2go"
	"errors"
)

type gitSeekretConfig struct {
	gitConfig *git.Config
	version int32
	rulespath string
	rulesenabled string
}

var gitSeekretDefaultConfig = &gitSeekretConfig{
	gitConfig: nil,
	version: 1,
	rulespath: s.DefaultRulesPath(),
	rulesenabled: "rulesenabled",
}

var gitSeekretCurrentConfig *gitSeekretConfig

func NewGitSeekretConfig() *gitSeekretConfig {
	return &gitSeekretConfig{}
}

func (gsc *gitSeekretConfig) LoadConfig(global bool) error {
	if global {
		configFile, err := git.ConfigFindGlobal()
		if err != nil {
			return err
		}

		gsc.gitConfig, err = git.OpenOndisk(nil, configFile)
		if err != nil {
			return err
		}
	} else {
		repo, err := openGitRepoLocal(".")
		if err != nil {
			return err
		}

		gsc.gitConfig, err = repo.Config()
		if err != nil {
			return err
		}
	}

	var err error

	if gsc.version, err =  gsc.gitConfig.LookupInt32("gitseekret.version") ; err != nil {
		gsc.version = gitSeekretDefaultConfig.version
	}

	if gsc.rulespath, _ =  gsc.gitConfig.LookupString("gitseekret.rulespath") ; err != nil {
		gsc.rulespath = gitSeekretDefaultConfig.rulespath
	}
	if gsc.rulesenabled, _ =  gsc.gitConfig.LookupString("gitseekret.rulesenabled") ; err != nil {
		gsc.rulesenabled = gitSeekretDefaultConfig.rulesenabled
	}

	return nil
}

func (gsc *gitSeekretConfig) SaveConfig() error {
	var err error

	if gsc.gitConfig == nil {
		return errors.New("invalid gitConfig")
	}

	err = gsc.gitConfig.SetInt32("gitseekret.version", gsc.version)
	if err != nil {
		return err
	}

	err = gsc.gitConfig.SetString("gitseekret.rulespath", gsc.rulespath)
	if err != nil {
		return err
	}

	err = gsc.gitConfig.SetString("gitseekret.rulesenabled", gsc.rulesenabled)
	if err != nil {
		return err
	}

	return nil
}

func (gsc *gitSeekretConfig) ShowConfig() {
	fmt.Println("Config:")
	fmt.Printf("\tgitseekret.version = %d\n", gsc.version)
	fmt.Printf("\tgitseekret.rulespath = %s\n", gsc.rulespath)
	fmt.Printf("\tgitseekret.rulesenabled = %s\n", gsc.rulesenabled)
}

func openGitRepoLocal(source string) (*git.Repository, error) {
	var repo *git.Repository
	var err error

	for {
		source, _ = filepath.Abs(source)
		repo, err = git.OpenRepository(source)
		if err == nil {
			break
		}
		
		if git.IsErrorClass(err, git.ErrClassOs) {
			return nil, err
		}
			
		if source == "/" {
			return nil, err
		}
		source = source + "/.."
	}
	return repo, nil
}