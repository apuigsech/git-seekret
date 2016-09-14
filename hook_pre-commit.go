package main

import (
	"fmt"
	"github.com/apuigsech/seekret-source-git"
	"runtime"
)

func HookPreCommitEnable(args []string) (string, error) {
	return "git seekret hook --run pre-commit\n", nil
}

func HookPreCommitDisable(args []string) error {
	return nil
}

func HookPreCommitRun(args []string) error {
	options := map[string]interface{}{
		"commit": false,
		"staged": true,
	}

	err := gs.seekret.LoadObjects(sourcegit.SourceTypeGit, ".", options)
	if err != nil {
		return err
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

	if len(listSecrets) > 0 {
		return fmt.Errorf("commit cannot proceed")
	}

	return nil
}
