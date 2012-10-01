package main

import (
	"fmt"
	"os"

	"github.com/sbinet/go-commander"
	"github.com/sbinet/go-flag"
)

func run_list_projects(cmd *commander.Command, args []string) {
	//n := cmd.Name()
	//fmt.Printf("%s:  args: %v\n", n, args)
	//fmt.Printf("%s: flags: %v\n", n, cmd.Flag.NArg())

	msg, err := g_ami.Execute("ListProject")
	if err != nil {
		fmt.Printf("**err** %v\n", err)
		os.Exit(1)
	}
	
	// fmt.Printf("msg: [%v]\n", msg.XMLName)
	// fmt.Printf("command:  %q\n", msg.Cmd)
	// fmt.Printf("cmd-args: %v\n", msg.CmdArgs)
	// fmt.Printf("status:   %q %v\n", msg.CmdStatus, msg.Status())
	// fmt.Printf("exectime: %vs\n", msg.ExecTime)
	// fmt.Printf("rows:     %v\n", len(msg.Result.Rows))
	for _,v := range msg.Result.Rows {
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
