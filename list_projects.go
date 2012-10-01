package main

import (
	"fmt"
	"os"

	"github.com/sbinet/go-commander"
	"github.com/sbinet/go-flag"
)

func run_list_projects(cmd *commander.Command, args []string) {

	msg, err := g_ami.Execute("ListProject")
	if err != nil {
		fmt.Printf("**err** %v\n", err)
		os.Exit(1)
	}

	for _, v := range msg.Result.Rows {
		m := v.Value()
		fmt.Printf("%s (descr=%q)\n", m["name"], m["description"])
	}
}

func ami_make_list_projects_cmd() *commander.Command {
	cmd := &commander.Command{
		Run:       run_list_projects,
		UsageLine: "list-projects",
		Short:     "List projects",
		Long: `
list-projects lists the projects in AMI.

ex:
 $ go-ami list-projects
`,
		Flag: *flag.NewFlagSet("ami-list-projects", flag.ExitOnError),
	}
	return cmd
}
