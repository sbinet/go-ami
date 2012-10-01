package main

import (
	"fmt"
	"os"

	"github.com/sbinet/go-commander"
	"github.com/sbinet/go-flag"
)

func run_any_cmd(cmd *commander.Command, args []string) {
	//n := cmd.Name()
	//fmt.Printf("%s:  args: %v\n", n, args)
	//fmt.Printf("%s: flags: %v\n", n, cmd.Flag.NArg())

	msg, err := g_ami.Execute(args...)
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
	for i, row := range msg.Result.Rows {
		fmt.Printf("row=%d\n", i)
		m := row.Value()
		for k, v := range m {
			fmt.Printf("  -> %s=%v\n", k, v)
		}
	}
}

func ami_make_cmd_cmd() *commander.Command {
	cmd := &commander.Command{
		Run:       run_any_cmd,
		UsageLine: "cmd <AmiCommand> <AmiArg0> <AmiArg1>...",
		Short:     "run an arbitrary AMI command",
		Long: `
cmd sends an arbitrary AMI command to the AMI server.

ex:
 $ go-ami cmd TCGetPackageInfo fullPackageName="/External/pyAMI" processingStep="production" project="TagCollector" repositoryName="AtlasOfflineRepository"
`,
		Flag: *flag.NewFlagSet("ami-cmd", flag.ExitOnError),
	}
	return cmd
}

// EOF
