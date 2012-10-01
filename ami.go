package main

import (
	"fmt"
	"os"

	"github.com/sbinet/go-ami/pkg/ami"
	"github.com/sbinet/go-commander"
	"github.com/sbinet/go-flag"
)

var g_ami *ami.Client
var g_cmd *commander.Commander

func main() {

	g_cmd = &commander.Commander{
		Name: os.Args[0],
		Commands: []*commander.Command{
			ami_make_cmd_cmd(),
			ami_make_list_runs_cmd(),
			ami_make_list_projects_cmd(),
			ami_make_setup_auth_cmd(),
		},
		Flag: flag.NewFlagSet("ami", flag.ExitOnError),
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
	if err != nil {
		fmt.Printf("**error** could not create ami.Client: %v\n", err)
		os.Exit(1)
	}
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
	err = g_cmd.Run(args)
	if err != nil {
		fmt.Printf("**error** %v\n", err)
		os.Exit(1)
	}
}

// EOF
