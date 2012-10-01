package main

import (
	"fmt"
	"os"

	"github.com/sbinet/go-ami/pkg/ami"
	"github.com/sbinet/go-commander"
	"github.com/sbinet/go-flag"
)

func run_list_runs(cmd *commander.Command, args []string) {

	// n := cmd.Name()
	// fmt.Printf("%s:  args: %v\n", n, args)
	// fmt.Printf("%s: flags: %v\n", n, cmd.Flag.NArg())
	year := cmd.Flag.Lookup("year").Value.Get().(int)
	if year > 2000 {
		year %= 1000
	}

	periods := cmd.Flag.Lookup("periods").Value.Get().(string)
	// fmt.Printf("%s: year=%v\n", n, year)

	runs, err := ami.GetRuns(g_ami, periods, year)
	if err != nil {
		fmt.Printf("**err** %v\n", err)
		os.Exit(1)
	}

	for _, run := range runs {
		fmt.Printf("%d\n", run)
	}
}

func ami_make_list_runs_cmd() *commander.Command {
	cmd := &commander.Command{
		Run:       run_list_runs,
		UsageLine: "list-runs -year YEAR -period PERIOD",
		Short:     "List runs in a data period for a given year",
		Long: `
list-runs lists the runs in a data period for a given year.

ex:
 $ go-ami list-runs -year 2012 -period M1
`,
		Flag: *flag.NewFlagSet("ami-list-runs", flag.ExitOnError),
	}
	cmd.Flag.Int("year", 2012, "year for the data period")
	cmd.Flag.String("periods", "", "comma-separated list of period name(s) e.g. M1,M2")
	return cmd
}
