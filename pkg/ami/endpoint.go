package ami

var EndPointType = "main"

func EndPoint() string {

    if EndPointType == "replica" {
        return "https://atlas-ami.cern.ch:8443/AMI/services/AMIWebService"
    }
    return "https://ami.in2p3.fr/AMI/services/AMIWebService"
}

// EOF
