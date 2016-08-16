package main

import (
	"fmt"
	"os"
	"github.com/libgit2/git2go"
	"github.com/urfave/cli"
)

type GitSeekretHookHandler struct {
	Enable func(args []string) (string, error)
	Disable func(args []string) (error)
	Run func(args []string) (error)
}

var listGitSeekretHookHandler map[string]GitSeekretHookHandler = map[string]GitSeekretHookHandler{
	"pre-commit": GitSeekretHookHandler{
		Enable: HookPreCommitEnable,
		Disable: HookPreCommitDisable,
		Run: HookPreCommitRun,
	},
}


func GitSeekretHook(c *cli.Context) error {
	// TODO: Implement also support for --global
	err := gs.LoadConfig(git.ConfigLevelLocal, true)	
	if git.IsErrorClass(err, git.ErrClassConfig) {
		return fmt.Errorf("Config not initialised - Try: 'git-seekret config --init'")
	}
	if err != nil {
		return err
	}	

	enable := c.String("enable")
	disable := c.String("disable")

	if enable != "" {
		err := GitSeekretHookEnable(enable)
		if err != nil {
			return err
		}
	}

	if disable != "" {
		err := GitSeekretHookDisable(disable)
		if err != nil {
			return err
		}
	}

	if c.Bool("enable-all") {
		for name,_ := range listGitSeekretHookHandler {
			err := GitSeekretHookEnable(name)
			if err != nil {
				return err
			}
		}
	}

	if c.Bool("disable-all") {
		for name,_ := range listGitSeekretHookHandler {
			err := GitSeekretHookDisable(name)
			if err != nil {
				return err
			}
		}
	}	

	run := c.String("run")

	if run != "" {
		handler, ok := listGitSeekretHookHandler[run]
		if ok && handler.Run != nil {
			err := handler.Run(nil)
			if err != nil {
				return err
			}
		}		
	}

	return nil
}

func GitSeekretHookEnable(name string) (error) {
	var script string
	var err error

	handler, ok := listGitSeekretHookHandler[name]
	if ok && handler.Enable != nil {
		script,err = handler.Enable(nil)
		if err != nil {
			return err
		}
	}

	_ = script

	hookfile := fmt.Sprintf("%s/hooks/%s", gs.repo, name)

	if _, err := os.Stat(hookfile); err == nil {
		hookfile_old := fmt.Sprintf("%s/hooks/%s.old", gs.repo, name)
		err = os.Rename(hookfile, hookfile_old)
		if err != nil {
			return err
		}
	}

	fh, err := os.Create(hookfile)
	if err != nil {
		return err
	}
	defer fh.Close()

	fh.WriteString("#!/usr/bin/env bash\n\n")
	fh.WriteString(script)
	fh.Close()

	err = os.Chmod(hookfile, 0755)
	if err != nil {
		return err
	}

	return nil
}

func GitSeekretHookDisable(name string) (error) {
	handler, ok := listGitSeekretHookHandler[name]
	if ok && handler.Disable != nil {
		err := handler.Disable(nil)
		if err != nil {
			return err
		}
	}

	hookfile := fmt.Sprintf("%s/hooks/%s", gs.repo, name)

	err := os.Remove(hookfile)
	if err != nil {
		return err
	}

	return nil
}

