package ami

import (
	"fmt"
)

var EndPointType = "main"

func EndPoint() string {
	switch EndPointType {
	case "replica":
		return "https://atlas-ami.cern.ch:8443"
	case "main":
		return "https://ami.in2p3.fr:8443"
	default:
		fmt.Printf("**error** invalid ami.EndPointType: %q\n", EndPointType)
	}
	panic("unreachable")
}

func XslUrl() string {
	switch EndPointType {
	case "replica":
		return "https://atlas-ami.cern.ch:8443/AMI/AMI/xsl/"
	case "main":
		return "https://ami.in2p3.fr/AMI/AMI/xsl/"
	default:
		fmt.Printf("**error** invalid ami.EndPointType: %q\n", EndPointType)
	}
	panic("unreachable")
}

// EOF
