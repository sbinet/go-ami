package ami

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Client struct {
	verbose  bool
	vformat  string
	config   Config
	cert     *tls.Certificate
	client   *http.Client
	nqueries int // number of possible concurrent queries
}

func NewClient(verbose bool, format string, nqueries int) *Client {
	c := &Client{
		verbose:  verbose,
		vformat:  format,
		config:   NewConfig(),
		client:   nil,
		nqueries: nqueries,
	}
	nqueries = 5
	if c.nqueries < 1 {
		fmt.Printf("ami: nqueries too low (==%v). setting to %v\n",
			c.nqueries, nqueries)
		c.nqueries = nqueries
	}
	if c.nqueries > 10 {
		fmt.Printf("ami: nqueries too high (==%v). setting to %v\n",
			c.nqueries, nqueries)
		c.nqueries = nqueries
	}

	err := c.authenticate()
	if err != nil {
		fmt.Printf("**err** %v\n", err)
		return nil
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{*c.cert},
			InsecureSkipVerify: true,
		},
	}
	c.client = &http.Client{Transport: tr}

	return c
}

func (c *Client) Execute(args ...string) (*Message, error) {

	var err error

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
			arg = arg[0:idx]
		}
		amiargs = append(amiargs, fmt.Sprintf("-%s=%s", arg, val))
	}
	cmd.Add("Command", strings.Join(amiargs, " "))
	if c.verbose {
		fmt.Printf("==> %v\n", amiargs)
	}
	path := "/AMI/servlet/net.hep.atlas.Database.Bookkeeping.AMI.Servlet.Command"
	rawurl := NewSvcLocator().ServiceAddress() + path + "?" + cmd.Encode()
	if c.verbose {
		fmt.Printf("url: %v\n", rawurl)
	}

	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
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

	cert_fname := filepath.Join(os.ExpandEnv(CertDir), "usercert.pem")
	key_fname := filepath.Join(os.ExpandEnv(CertDir), "userkey.pem")

	//FIXME: 
	// X509_USER_PROXY doesn't have a format tls.LoadX509KeyPair understands
	// if os.Getenv("X509_USER_PROXY") != "" {
	// 	cert_fname = os.Getenv("X509_USER_PROXY")
	// 	key_fname = cert_fname
	// }

	cert, err := tls.LoadX509KeyPair(cert_fname, key_fname)
	if err != nil {
		user_cert, user_key, err := LoadCert(cert_fname, key_fname)
		if err != nil {
			fmt.Printf("ami: error while trying to load certificate. try running:\n  $ go-ami setup-auth\n")
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

// EOF
