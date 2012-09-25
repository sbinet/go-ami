package ami

import (
	"fmt"
	"net/http"
	"net/url"
)

type ServiceLocator interface {
	ServiceAddress() string
	WebService(url string) (WebService, error)
}

type WebService interface {
	ExecCmd(req *http.Request) (*http.Response, error)
}

func NewSvcLocator() ServiceLocator {
	return &secureWebSvcLocator{}
}

type secureWebSvcLocator struct {}

func (svcloc *secureWebSvcLocator) ServiceAddress() string {
	return EndPoint()
}

func (svcloc *secureWebSvcLocator) WebService(rawurl string) (WebService,error) {
	if rawurl == "" {
		rawurl = svcloc.ServiceAddress()
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	wsvc := &soapWebSvc{
	url: u,
	client: new(http.Client),
	}
	return wsvc, nil
}

type soapWebSvc struct {
	url *url.URL
	client *http.Client
}

func (svc *soapWebSvc) ExecCmd(request *http.Request) (*http.Response, error) {
	fmt.Printf("==ExecCmd(%q)...\n", request)
	
	return nil, nil
}
// EOF
