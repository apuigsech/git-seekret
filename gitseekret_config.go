package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"strings"
	seekret "github.com/apuigsech/seekret/lib"
)

type gitSeekretConfig struct {
	version int32
	rulespath string
	rulesenabled string
	exceptionsfile string
}

func NewGitSeekretConfigInit() (*gitSeekretConfig) {
	gsc := &gitSeekretConfig{
		version: 1,
		rulespath: seekret.DefaultRulesPath(),
		rulesenabled: "",
		exceptionsfile: "",
	}
	return gsc
}

func NewGitSeekretConfigLoad(gitConfig *git.Config) (*gitSeekretConfig,error) {
	if gitConfig == nil {
		return nil,fmt.Errorf("invalid gitConfig")
	}

	version, err := gitConfig.LookupInt32("gitseekret.version")
	if err != nil {
		return nil,err
	}

	rulespath, err := gitConfig.LookupString("gitseekret.rulespath")
	if err != nil {
		return nil,err
	}

	rulesenabled, err := gitConfig.LookupString("gitseekret.rulesenabled")
	if err != nil {
		return nil,err
	}

	exceptionsfile, err := gitConfig.LookupString("gitseekret.exceptionsfile")
	if err != nil {
		return nil,err
	}


	gsc := &gitSeekretConfig{
		version: version,
		rulespath: rulespath,
		rulesenabled: rulesenabled,
		exceptionsfile: exceptionsfile,
	}

	return gsc,nil
}


func (gsc *gitSeekretConfig)Save(gitConfig *git.Config) (error) {
	if gitConfig == nil {
		return fmt.Errorf("invalid gitConfig")
	}

	err := gitConfig.SetInt32("gitseekret.version", gsc.version)
	if err != nil {
		return err
	}

	err = gitConfig.SetString("gitseekret.rulespath", gsc.rulespath)
	if err != nil {
		return err
	}

	err = gitConfig.SetString("gitseekret.rulesenabled", gsc.rulesenabled)
	if err != nil {
		return err
	}

	err = gitConfig.SetString("gitseekret.exceptionsfile", gsc.exceptionsfile)
	if err != nil {
		return err	
	}

	return nil
}

func (gsc *gitSeekretConfig)Run(s *seekret.Seekret) (error) {
	// TODO: Relative path from repo root.
	err := s.LoadRulesFromPath(gsc.rulespath, false)
	if err != nil {
		return err
	}

	for _,rule := range strings.Split(gsc.rulesenabled, ",") {
		s.EnableRule(rule)
	}

	// TODO: Relative path from repo root.
	if gsc.exceptionsfile != "" {
		err := s.LoadExceptionsFromFile(gsc.exceptionsfile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (gsc *gitSeekretConfig)BuildRulesEnabled(s *seekret.Seekret) {
	listRules := s.ListRules()
	rulesenabled := make([]string, 0, len(listRules))
	for _,rule := range gs.seekret.ListRules() {
		if rule.Enabled {
			rulesenabled = append(rulesenabled, rule.Name)
		}
	}
	gsc.rulesenabled = strings.Join(rulesenabled, ",")
}

