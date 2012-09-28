package ami

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type Client struct {
	verbose    bool
	vformat    string
	config     Config
	cert       *tls.Certificate
	client     *http.Client
	//svclocator ServiceLocator
}

func NewClient(verbose bool, format string) *Client {
	c := &Client{
		verbose: verbose,
		vformat: format,
	config: NewConfig(),
	//svclocator: NewSvcLocator(),
	client: nil,
	}
	return c
}

func (c *Client) Execute(args ...string) (*Message, error) {

	err := c.authenticate()
	if err != nil {
		fmt.Printf("errrr\n")
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
		    Certificates: []tls.Certificate {*c.cert},
		    InsecureSkipVerify: true,
		//ClientAuth:  tls.RequireAndVerifyClientCert,
		},
	}
	c.client = &http.Client{Transport: tr}

	cmd := url.Values{}
	amiargs := []string{args[0]}
	for _, arg := range args[1:] {
		val := ""
		if strings.HasPrefix(arg, "-") {
			arg = arg[1:]
			if strings.HasPrefix(arg, "-") {
				arg = arg[1:]
			}
		}
		if idx := strings.Index(arg, "="); idx > -1 {
			val = arg[idx+1:]
			//val = strings.Replace(val, "=", "\=", -1)
			arg = arg[0:idx]
		}
		amiargs = append(amiargs, fmt.Sprintf("-%s=%s", arg,val))
	}
	cmd.Add("Command", strings.Join(amiargs, " "))
	fmt.Printf("==> %v\n", amiargs)
	path := "/AMI/servlet/net.hep.atlas.Database.Bookkeeping.AMI.Servlet.Command"
	rawurl := NewSvcLocator().ServiceAddress()+path+"?"+cmd.Encode()
	fmt.Printf("url: %v\n", rawurl)

    req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
    if err!=nil {
		return nil, err
    }

	if resp.Status != "200 OK" {
		fmt.Println(resp.Status)
		return nil, fmt.Errorf("httperror: %v\n", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//fmt.Printf("==> %v\n", string(data))

	msg := Message{}
	err = xml.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (c *Client) authenticate() error {
	if c.cert != nil {
		return nil
	}

	cert_fname := ""
	key_fname := ""

	if os.Getenv("X509_USER_PROXY") != "" {
		cert_fname = os.Getenv("X509_USER_PROXY")
		key_fname = cert_fname
	} else {
		cert_fname = os.ExpandEnv("${HOME}/.config/go-ami/certs/usercert.pem")
		key_fname = os.ExpandEnv("${HOME}/.config/go-ami/certs/userkey.pem")
	}

	cert, err := tls.LoadX509KeyPair(cert_fname, key_fname)
	if err != nil {
		user_cert, user_key, err := load_cert(cert_fname, key_fname)
		if err != nil {
			return err
		}
		cert, err = tls.X509KeyPair(user_cert, user_key)
		if err != nil {
			return err
		}
	}
	c.cert = &cert
	return nil
}

func load_cert(cert_fname, key_fname string) (user_cert, user_key []byte, err error) {
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
