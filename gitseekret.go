package main

import (
	"fmt"
	seekret "github.com/apuigsech/seekret/lib"
	"github.com/libgit2/git2go"
)

const gitSeekretConfigVersion = 1

type gitSeekret struct {
	repo string
	seekret *seekret.Seekret

	configLevel git.ConfigLevel
	gitConfig *git.Config
	gitSeekretConfig *gitSeekretConfig
}


func NewGitSeekret(configLevel git.ConfigLevel, repo string) (*gitSeekret, error) {
	gs := &gitSeekret{
		configLevel: configLevel,
		repo: repo,
	}

	return gs,nil
}

func (gs *gitSeekret)InitConfig() (error) {
	var err error

	gs.seekret = seekret.NewSeekret()

	gs.gitConfig, err = openGitConfig(gs.configLevel, gs.repo)
	if err != nil {
		return err
	}
	defer gs.gitConfig.Free()

	gs.gitSeekretConfig = NewGitSeekretConfigInit()

	err = gs.gitSeekretConfig.Run(gs.seekret)
	if err != nil {
		return err
	}

	return nil
}


func (gs *gitSeekret)LoadConfig(run bool) (error) {
	var err error

	gs.seekret = seekret.NewSeekret()

	gs.gitConfig, err = openGitConfig(gs.configLevel, gs.repo)
	if err != nil {
		return err
	}
	defer gs.gitConfig.Free()

	gs.gitSeekretConfig,err = NewGitSeekretConfigLoad(gs.gitConfig)
	if err != nil {
		return err
	}

	if run {
		err = gs.gitSeekretConfig.Run(gs.seekret)
		if err != nil {
			return err
		}
	}

	return nil
}


func (gs *gitSeekret)SaveConfig() (error) {
	if gs.gitConfig == nil {
		return fmt.Errorf("git config not loaded")
	}

	
	gs.gitSeekretConfig.BuildRulesEnabled(gs.seekret)
	
	err := gs.gitSeekretConfig.Save(gs.gitConfig)
	if err != nil {
		return err
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


func (gs *gitSeekret)EnableRule(name string) (error) {
	return gs.seekret.EnableRule(name)
}

func (gs *gitSeekret)DisableRule(name string) (error) {
	return gs.seekret.DisableRule(name) 
}