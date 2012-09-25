package main

import (
	"fmt"

	"github.com/sbinet/go-commander"
	"github.com/sbinet/go-flag"
)

func run_list_projects(cmd *commander.Command, args []string) {
	n := cmd.Name()
	fmt.Printf("%s:  args: %v\n", n, args)
	fmt.Printf("%s: flags: %v\n", n, cmd.Flag.NArg())

	g_ami.Execute([]string{"ListProject"})
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
