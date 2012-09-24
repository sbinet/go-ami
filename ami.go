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
			ami_make_list_cmd(),
		},
	Flag: flag.NewFlagSet("ami", flag.ExitOnError),
	}
	//TODO: check the value *is* in the [main,replica] list
	g_cmd.Flag.String("server", "main", "set the server (main, replica)")

	err := g_cmd.Flag.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("**err** %v\n", err)
		os.Exit(1)
	}
	args := g_cmd.Flag.Args()

	fmt.Printf("%s: server=%v\n", g_cmd.Name, g_cmd.Flag.Lookup("server").Value)
	fmt.Printf("%s: args=%v\n", g_cmd.Name, args)
	err = g_cmd.Run(args)
	if err != nil {
		fmt.Printf("**err** %v\n", err)
		os.Exit(1)
	}
}

// EOF
