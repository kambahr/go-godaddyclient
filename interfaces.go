package gogadaddyclient

import (
	"fmt"
	"net"
)

// IGodaddyAPI is the GoaddyClient interface.
type IGodaddyAPI interface {
	CreateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result
	DeleteSubDomain(domainName string, subDomainName string) Result
}

func (g *GodaddyAPI) CreateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result {

	var res Result

	n := net.ParseIP(forwardIPAddr)
	if n == nil {
		res.Message = "invalid ip address"
		return res
	}

	url := fmt.Sprintf("%s/v1/domains/%s/records/A/%s", g.BaseURL, domainName, subDomainName)
	jsonStr := fmt.Sprintf(`[{"data": "%s","ttl": 600}]`, forwardIPAddr)

	res = g.execute("PUT", url, jsonStr)

	return res
}

func (g *GodaddyAPI) DeleteSubDomain(domainName string, subDomainName string) Result {

	var res Result

	url := fmt.Sprintf("%s/v1/domains/%s/records/A/%s", g.BaseURL, domainName, subDomainName)
	res = g.execute("DELETE", url, "")

	return res
}
