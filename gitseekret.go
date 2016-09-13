package main

import (
	"github.com/apuigsech/seekret"
	"github.com/libgit2/git2go"
)

const gitSeekretConfigVersion = 1

type gitSeekret struct {
	repo string
	seekret *seekret.Seekret
	config *gitSeekretConfig
}


func NewGitSeekret(repo string) (*gitSeekret, error) {
	var err error 

	repo, err = repoBasePath(repo)
	if err != nil {
		return nil,err
	}

	gs := &gitSeekret{
		repo: repo,
		seekret: seekret.NewSeekret(),
	}

	return gs,nil
}



func (gs *gitSeekret)EnableRule(name string) int {
	return gs.seekret.EnableRuleByRegexp(name)
}


func (gs *gitSeekret)DisableRule(name string) int {
	return gs.seekret.DisableRuleByRegexp(name)
}


func repoBasePath(repo string) (string, error) {
	r, err := git.OpenRepositoryExtended(repo, git.RepositoryOpenCrossFs, "")
	if err != nil {
		return "", err
	}

	path := r.Path()

	r.Free()

	return path,nil
}