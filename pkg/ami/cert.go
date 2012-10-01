package ami

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func LoadCert(cert_fname, key_fname string) (user_cert, user_key []byte, err error) {
	user_cert, err = ioutil.ReadFile(cert_fname)
	if err != nil {
		return
	}

	// decrypt key-file
	cmd := exec.Command("openssl", "rsa", "-in", key_fname)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	user_key, err = cmd.Output()
	if err != nil {
		return
	}

	return
}

// EOF
