package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
	"github.com/sbinet/go-ami/ami"
)

func path_exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func run_setup_auth(cmd *commander.Command, args []string) error {
	n := cmd.Name()
	// fmt.Printf("%s:  args: %v\n", n, args)
	// fmt.Printf("%s: flags: %v\n", n, cmd.Flag.NArg())

	dirname := os.ExpandEnv(ami.ConfigDir)
	if !path_exists(dirname) {
		err := os.MkdirAll(dirname, 0700)
		if err != nil {
			return err
		}
	}
	dirname = os.ExpandEnv(ami.CertDir)
	if !path_exists(dirname) {
		err := os.MkdirAll(dirname, 0700)
		if err != nil {
			return err
		}
	}

	cert_fname := cmd.Flag.Lookup("usercert").Value.Get().(string)
	if !path_exists(cert_fname) {
		fmt.Printf("%s: no such user certificate file [%s]\n", n, cert_fname)
	}
	key_fname := cmd.Flag.Lookup("userkey").Value.Get().(string)
	if !path_exists(key_fname) {
		fmt.Printf("%s: no such user key file [%s]\n", n, key_fname)
	}

	user_cert, user_key, err := ami.LoadCert(cert_fname, key_fname)
	if err != nil {
		return err
	}

	cert_fname = filepath.Join(dirname, "usercert.pem")
	err = ioutil.WriteFile(cert_fname, user_cert, 0600)
	if err != nil {
		return err
	}

	key_fname = filepath.Join(dirname, "userkey.pem")
	err = ioutil.WriteFile(key_fname, user_key, 0600)
	if err != nil {
		return err
	}

	return nil
}

func ami_make_setup_auth_cmd() *commander.Command {
	cmd := &commander.Command{
		Run:       run_setup_auth,
		UsageLine: "setup-auth -userkey userkey.pem -usercert usercert.pem",
		Short:     "Setup authentication mechanism for go-ami",
		Long: `
setup-auth setups go-ami for proper certificate authentication.

ex:
 $ go-ami setup-auth -userkey ~/.grid/userkey.pem -usercert ~/.grid/usercert.pem
`,
		Flag: *flag.NewFlagSet("ami-setup-auth", flag.ExitOnError),
	}
	cmd.Flag.String("userkey", "userkey.pem", "file holding the PEM user key")
	cmd.Flag.String("usercert", "usercert.pem", "file holding the PEM user certificate")
	return cmd
}

// EOF
