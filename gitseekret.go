package main

import (
	seekret "github.com/apuigsech/seekret/lib"
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



func (gs *gitSeekret)EnableRule(name string) (error) {
	return gs.seekret.EnableRule(name)
}


func (gs *gitSeekret)DisableRule(name string) (error) {
	return gs.seekret.DisableRule(name) 
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


/*
func main() {
	gs,err := NewGitSeekret(".")
	if err != nil {
		fmt.Println(err)
	}
	gs.LoadConfig(git.ConfigLevelLocal, true)
	fmt.Printf("%#v\n", gs)
	fmt.Printf("%#v\n", gs.config)
	fmt.Printf("%#v\n", gs.seekret.ListRules())
}
*/