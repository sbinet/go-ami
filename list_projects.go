package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func run_list_projects(cmd *commander.Command, args []string) {

	msg, err := g_ami.Execute("ListProject")
	if err != nil {
		fmt.Printf("**error** %v\n", err)
		os.Exit(1)
	}

	projs := make([]string, 0, len(msg.Result.Rows))
	for _, v := range msg.Result.Rows {
		m := v.Value()
		projs = append(projs,
			fmt.Sprintf("%s (descr=%q)", m["name"].(string), m["description"].(string)),
		)
	}
	sort.Strings(projs)
	for _, proj := range projs {
		fmt.Printf("%s\n", proj)
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
