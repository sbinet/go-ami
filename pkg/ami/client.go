package ami

import (
	//"bufio"
	"fmt"
	"net/http"
	//"os"
	"strings"
)

type Client struct {
	verbose    bool
	vformat    string
	config     Config
	svclocator ServiceLocator
}

func NewClient(verbose bool, format string) *Client {
	c := &Client{
		verbose: verbose,
		vformat: format,
	config: NewConfig(),
	svclocator: NewSvcLocator(),
	}
	return c
}

func (c *Client) Execute(args ...string) (*http.Response, error) {
	
	amiargs := []string{"-command="+args[0]}
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
	fmt.Printf("==> %v\n", amiargs)

	svc, err := c.svclocator.WebService("")
	if err != nil {
		return nil, err
	}

    soapRequestContent := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?><SOAP-ENV:En ... elope>"
    req, err := http.NewRequest("POST", c.svclocator.ServiceAddress(), 
		strings.NewReader(soapRequestContent))
    req.Header.Set("SOAPAction", "anAction")
    req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
    req.Header.Set("Content-Length", fmt.Sprintf("%d", len(soapRequestContent)))

	resp, err := svc.ExecCmd(req)
    if err!=nil {
		return nil, err
    }
	/*
    r := bufio.NewReader(resp.Body)
    line, err := r.ReadString('\n')
    for err == nil {
        fmt.Print(line)
        line, err = r.ReadString('\n')
    }
    if err != os.EOF {
        fmt.Println("Error :")
        fmt.Println(err)
    }
    //resp.Body.Close()
	 */
	return resp, err
}

// EOF
