package main

import (
	"fmt"
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

	secrets, err := gs.RunCheck(options)
	if err != nil {
		return err
	}
	if secrets != 0 {
		return fmt.Errorf("commit cannot proceed")
	}

	return nil
}
