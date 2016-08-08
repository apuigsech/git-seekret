package main

import(
	"fmt"
	"os"
	"bufio"
	"github.com/urfave/cli"
)

func GitSeekretRules(c *cli.Context) error {
	enable := c.String("enable")
	disable := c.String("disable")

	if enable != "" {
		s.EnableRule(enable)
	}

	if disable != "" {
		s.DisableRule(disable)
	}

	fmt.Println("List of rules:")
	for _, r := range s.ListRules() {
		status := " "
		if r.Enabled {
			status = "x"
		}
		fmt.Printf("\t[%s] %s\n", status, r.Name)
	}
	return nil
}

func loadEnabledRules() error {
	if _, err := os.Stat(gitSeekretCurrentConfig.rulesenabled); os.IsNotExist(err) {
		return nil
	}

	fh, err := os.Open(gitSeekretCurrentConfig.rulesenabled)
	if err != nil {
    	return err
  	}	
  	defer fh.Close()

  	scanner := bufio.NewScanner(fh)
  	for scanner.Scan() {
  		rule := scanner.Text()
  		s.EnableRule(rule)
  	}
  	return nil
}

func saveEnabledRules() error {
	fh, err := os.Create(gitSeekretCurrentConfig.rulesenabled)
	if err != nil {
    	return err
  	}
	defer fh.Close()

	w := bufio.NewWriter(fh)
	for _,r := range s.ListRules() {
		if r.Enabled {
			fmt.Fprintln(w, r.Name)
		}
	}
	return w.Flush()
}