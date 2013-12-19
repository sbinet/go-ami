package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func run_any_cmd(cmd *commander.Command, args []string) error {
	//n := cmd.Name()
	//fmt.Printf("%s:  args: %v\n", n, args)
	//fmt.Printf("%s: flags: %v\n", n, cmd.Flag.NArg())

	msg, err := g_ami.Execute(args...)
	if err != nil {
		return err
	}

	for i, row := range msg.Result.Rows {
		fmt.Printf("row=%d\n", i)
		m := row.Value()
		for k, v := range m {
			fmt.Printf("  -> %s=%v\n", k, v)
		}
	}

	return nil
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

row=0
  -> computingDescription=0
  -> created=2012-08-31 15:41:53 +0000 UTC
  -> physicsDescription=
  -> archive=0
  -> groupOriginOfPackage=AtlasRelease
  -> fullPackageName=/External/pyAMI
  -> comment=
  -> packageName=pyAMI
  -> repositoryName=AtlasOfflineRepository
  -> path=/External/
  -> softwareDomain=
  -> defaultType=leaf
row=1
 [...]
`,
		Flag: *flag.NewFlagSet("ami-cmd", flag.ExitOnError),
	}
	return cmd
}

// EOF
