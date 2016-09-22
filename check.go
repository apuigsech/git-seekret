package main

import (
	"fmt"
	"github.com/apuigsech/seekret-source-git"
	"runtime"
)

func (gs *gitSeekret) RunCheck(options map[string]interface{}) (int, error) {

	err := gs.seekret.LoadObjects(sourcegit.SourceTypeGit, gs.repo, options)
	if err != nil {
		return 0, err
	}

	gs.seekret.Inspect(runtime.NumCPU())

	listSecrets := gs.seekret.ListSecrets()
	fmt.Printf("Found Secrets: %d\n", len(listSecrets))
	for _, s := range listSecrets {
		fmt.Printf("\t%s:%d\n", s.Object.Name, s.Nline)
		fmt.Printf("\t\t- Metadata:\n")
		for k, v := range s.Object.Metadata {
			fmt.Printf("\t\t  %s: %s\n", k, v)
		}
		fmt.Printf("\t\t- Rule:\n\t\t  %s\n", s.Rule.Name)

		fmt.Printf("\t\t- Content:\n\t\t  %s\n", s.Line)
	}

	return len(listSecrets), nil

}
