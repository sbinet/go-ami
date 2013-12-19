package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/go-ami/ami"
)

func run_list_datasets(cmd *commander.Command, args []string) error {
	pattern := "%"
	if len(args) > 0 {
		pattern = args[0]
	}

	parent_type := cmd.Flag.Lookup("parent-type").Value.Get().(string)
	order := cmd.Flag.Lookup("order").Value.Get().(string)
	level := cmd.Flag.Lookup("limit").Value.Get().(int)
	fields := cmd.Flag.Lookup("fields").Value.Get().(string)
	flatten := false //FIXME
	show_archived := cmd.Flag.Lookup("show-archived").Value.Get().(bool)
	opts := make(map[string]interface{})

	datasets, err := ami.GetDatasets(g_ami, pattern, parent_type, order, level,
		fields, flatten, show_archived, opts)
	if err != nil {
		return err
	}

	for _, v := range datasets {
		fmt.Printf("%s\n", v)
	}

	return nil
}

func ami_make_list_datasets_cmd() *commander.Command {
	cmd := &commander.Command{
		Run:       run_list_datasets,
		UsageLine: "list-datasets [options] a-pattern",
		Short:     "List datasets matching a given pattern",
		Long: `
list-datasets lists datasets matching a given pattern.

ex:
 $ go-ami list-datasets --project mc12_8TeV --type NTUP_TAUMEDIUM %
mc12_8TeV.105200.McAtNloJimmy_CT10_ttbar_LeptonFilter.merge.NTUP_TAUMEDIUM.e1193_s1469_s1470_r3542_r3549_p1011
mc12_8TeV.105204.McAtNloJimmy_AUET2CT10_ttbar_allhad.merge.NTUP_TAUMEDIUM.e1305_s1469_s1470_r3542_r3549_p1011
mc12_8TeV.105334.HerwigVBFH120tautaulh.merge.NTUP_TAUMEDIUM.e825_s1310_s1300_r3617_p1011
mc12_8TeV.105338.HerwigVBFH120tautauhh.merge.NTUP_TAUMEDIUM.e825_s1310_s1300_r3618_p1011
[...]
`,
		Flag: *flag.NewFlagSet("ami-list-datasets", flag.ExitOnError),
	}
	cmd.Flag.String("in-container", "", "")
	cmd.Flag.String("physics-comment", "", "")
	cmd.Flag.Int("run", -1, "")
	cmd.Flag.String("trash-trigger", "", "")
	cmd.Flag.String("trigger-config", "", "")
	cmd.Flag.String("period", "", "")
	cmd.Flag.Int("dataset-number", -1, "")
	cmd.Flag.String("requested-by", "", "")
	cmd.Flag.String("beam", "", "")
	cmd.Flag.String("creation-comment", "", "")
	cmd.Flag.String("prodsys-status", "", "")
	cmd.Flag.Int("nfiles", -1, "")
	cmd.Flag.String("conditions-tag", "", "")
	cmd.Flag.String("principal-physics-group", "", "")
	cmd.Flag.String("job-config", "", "")
	cmd.Flag.String("physics-process", "", "")
	cmd.Flag.String("modified-by", "", "")
	cmd.Flag.String("name", "", "")
	cmd.Flag.String("physics-category", "", "")
	cmd.Flag.String("created", "", "")
	cmd.Flag.String("geometry", "", "")
	cmd.Flag.String("atlas-release", "", "")
	cmd.Flag.String("responsible", "", "")
	cmd.Flag.String("created-by", "", "")
	cmd.Flag.String("transformation", "", "")
	cmd.Flag.String("project", "", "")
	cmd.Flag.String("trash-annotation", "", "")
	cmd.Flag.String("version", "", "")
	cmd.Flag.String("prod-step", "", "")
	cmd.Flag.String("physics-short", "", "")
	cmd.Flag.String("stream", "", "")
	cmd.Flag.String("type", "", "")
	cmd.Flag.String("events", "", "")
	cmd.Flag.String("ami-status", "VALID", "")
	cmd.Flag.String("history", "", "")
	cmd.Flag.String("files.guid", "", "")
	cmd.Flag.String("files.lfn", "", "")
	cmd.Flag.Int("files.events", -1, "")
	cmd.Flag.Int("files.size", -1, "")
	cmd.Flag.String("order", "", "order results by this field")
	cmd.Flag.Int("limit", -1, "limit number of results")
	cmd.Flag.String("fields", "", "extra fields (comma-separated) to display in output")
	cmd.Flag.Bool("show-archived", false, "search in archived catalogues as well")
	cmd.Flag.String("parent-type", "", "")
	return cmd
}
