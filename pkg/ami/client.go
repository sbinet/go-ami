package ami

type Client struct {
	verbose bool
	vformat string
	config  Config
	svclocator ServiceLocator
}

func NewClient(verbose bool, format string) *Client {
	c := &Client{
	verbose: verbose,
	vformat: format,
	}
	return c
}

// EOF
