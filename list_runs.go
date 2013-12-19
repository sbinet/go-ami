package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/go-ami/ami"
)

func run_list_runs(cmd *commander.Command, args []string) error {

	year := cmd.Flag.Lookup("year").Value.Get().(int)
	if year > 2000 {
		year %= 1000
	}

	periods := cmd.Flag.Lookup("periods").Value.Get().(string)
	// fmt.Printf("%s: year=%v\n", n, year)

	runs, err := ami.GetRuns(g_ami, periods, year)
	if err != nil {
		return err
	}

	for _, run := range runs {
		fmt.Printf("%d\n", run)
	}

	return nil
}

func ami_make_list_runs_cmd() *commander.Command {
	cmd := &commander.Command{
		Run:       run_list_runs,
		UsageLine: "list-runs -year YEAR -period PERIOD",
		Short:     "List runs in a data period for a given year",
		Long: `
list-runs lists the runs in a data period for a given year.

ex:
 $ go-ami list-runs -year 2012 -periods M1,M2
`,
		Flag: *flag.NewFlagSet("ami-list-runs", flag.ExitOnError),
	}
	cmd.Flag.Int("year", 2012, "year for the data period")
	cmd.Flag.String("periods", "", "comma-separated list of period name(s) e.g. M1,M2")
	return cmd
}
