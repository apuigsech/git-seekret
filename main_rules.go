package main

import(
	"fmt"
	"os"
	"bufio"
	"github.com/urfave/cli"
)

func loadRules() error {
	s.LoadRulesFromPath(os.Getenv("SEEKRET_RULES_PATH"), false)

	if _, err := os.Stat(GitSeekretEnabledRulesFile); os.IsNotExist(err) {
		return nil
	}

	fh, err := os.Open(GitSeekretEnabledRulesFile)
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

func saveRules() error {
	fh, err := os.Create(GitSeekretEnabledRulesFile)
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

func GitSeekretRulesList(c *cli.Context) error {
	listRules := s.ListRules()
	fmt.Println("List of rules:")
	for _, r := range listRules {
		status := " "
		if r.Enabled {
			status = "x"
		}
		fmt.Printf("\t[%s] %s\n", status, r.Name)
	}
	return nil
}

func GitSeekretRulesEnable(c *cli.Context) error {
	rule := c.Args().Get(0)
	if rule == ""  {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	s.EnableRule(rule)

	return nil
}

func GitSeekretRulesDisable(c *cli.Context) error {
	rule := c.Args().Get(0)
	if rule == ""  {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	s.DisableRule(rule)

	return nil
}