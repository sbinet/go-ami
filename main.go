package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/go-ami/ami"
)

var g_ami *ami.Client
var g_cmd *commander.Command

func main() {

	g_cmd = &commander.Command{
		UsageLine: "go-ami",
		Subcommands: []*commander.Command{
			ami_make_cmd_cmd(),
			ami_make_list_datasets_cmd(),
			ami_make_list_runs_cmd(),
			ami_make_list_projects_cmd(),
			ami_make_setup_auth_cmd(),
		},
		Flag: *flag.NewFlagSet("ami", flag.ExitOnError),
	}
	g_cmd.Flag.Bool("verbose", false, "show verbose output")
	g_cmd.Flag.Bool("debug", false, "show a stack trace")
	g_cmd.Flag.String("format", "text", "format of verbose output")
	g_cmd.Flag.Int("n", 5, "number of concurrent queries")

	//TODO: check the value *is* in the [main,replica] list via a special
	//      flag.Value implementation ?
	g_cmd.Flag.String("server", "main", "set the server (main, replica)")

	err := g_cmd.Flag.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("**error** %v\n", err)
		os.Exit(1)
	}
	g_ami, err = ami.NewClient(
		g_cmd.Flag.Lookup("verbose").Value.Get().(bool),
		g_cmd.Flag.Lookup("format").Value.Get().(string),
		g_cmd.Flag.Lookup("n").Value.Get().(int),
	)
	server := g_cmd.Flag.Lookup("server").Value.Get().(string)
	if server != "main" && server != "replica" {
		fmt.Printf("**error**. server has to be either 'main' or 'replica' (got: %q)\n", server)
		os.Exit(1)
	}
	ami.EndPointType = server

	args := g_cmd.Flag.Args()

	if g_cmd.Flag.Lookup("verbose").Value.Get().(bool) {
		fmt.Printf("%s: server=%v\n", g_cmd.Name, g_cmd.Flag.Lookup("server").Value)
		fmt.Printf("%s: args=%v\n", g_cmd.Name, args)
	}
	if len(args) > 0 && args[0] != "help" && args[0] != "setup-auth" {
		if err != nil {
			fmt.Printf("**error** could not create ami.Client: %v\n", err)
			if err == ami.ErrAuth {
				fmt.Printf("**error** try running:\n  go-ami setup-auth\n")
			}
			os.Exit(1)
		}
	}
	err = g_cmd.Dispatch(args)
	if err != nil {
		fmt.Printf("**error** %v\n", err)
		os.Exit(1)
	}
}

// EOF
