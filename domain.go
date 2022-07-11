package godaddyclient

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

type Domain struct {
	api IGodaddyAPI
}

// IDomain is the GoaddyClient interface.
type IDomain interface {
	CreateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result
	UpdateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result
	DeleteSubDomain(domainName string, subDomainName string) Result
	UpdateDomain(domainName string, forwardIPAddr string) Result
	DomainAvailable(domainName string) Result

	// CreateDNSRecord creates a new record.
	// domainName: top-level domain name,
	// recordType: type of DNS record,
	// name: name of the redord,
	// value: the value of the record,
	// ttl: the amount of time before the record is udpated (refrered)
	CreateDNSRecord(domainName string, recordType dnsRecordType, name string, value string, ttl int) Result
	UpdateDNSRecord(domainName string, recordType dnsRecordType, name string, value string, ttl int) Result
	DeleteDNSRecord(domainName string, recordType dnsRecordType, name string) Result
}

func getRecordType(rt dnsRecordType) (string, error) {

	switch rt {
	case a:
		return "A", nil
	case aaaa:
		return "AAAAA", nil
	case cname:
		return "CName", nil
	case mx:
		return "MX", nil
	case txt:
		return "TXT", nil
	case srv:
		return "SRV", nil
	case caa:
		return "CAA", nil
	case ns:
		return "NS", nil
	}

	return "", errors.New("invalid dns record type")
}

func (d *Domain) DomainAvailable(domainName string) Result {

	var res Result

	url := fmt.Sprintf("/domains/available?domain=%s&checkType=fast&forTransfer=false", domainName)
	res = d.api.HTTPExec(GET, url, "")

	return res
}

func (d *Domain) UpdateDomain(domainName string, forwardIPAddr string) Result {

	var res Result

	n := net.ParseIP(forwardIPAddr)
	if n == nil {
		res.Message = "invalid ip address"
		return res
	}
	subDomainName := "@"
	url := fmt.Sprintf("/domains/%s/records/A/%s", domainName, subDomainName)
	jsonStr := fmt.Sprintf(`[{"data": "%s","ttl": 600}]`, forwardIPAddr)

	res = d.api.HTTPExec(PUT, url, jsonStr)

	return res
}

func (d *Domain) UpdateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result {

	var res Result

	n := net.ParseIP(forwardIPAddr)
	if n == nil {
		res.Message = "invalid ip address"
		return res
	}

	url := fmt.Sprintf("/domains/%s/records/A/%s", domainName, subDomainName)
	jsonStr := fmt.Sprintf(`[{"data": "%s","ttl": 600}]`, forwardIPAddr)

	res = d.api.HTTPExec(PUT, url, jsonStr)

	return res
}

func (d *Domain) CreateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result {

	var res Result

	n := net.ParseIP(forwardIPAddr)
	if n == nil {
		res.Message = "invalid ip address"
		return res
	}

	url := fmt.Sprintf("/domains/%s/records/A/%s", domainName, subDomainName)
	jsonStr := fmt.Sprintf(`[{"data": "%s","ttl": 600}]`, forwardIPAddr)

	res = d.api.HTTPExec(PUT, url, jsonStr)

	return res
}
func (d *Domain) DeleteSubDomain(domainName string, subDomainName string) Result {

	var res Result

	url := fmt.Sprintf("/domains/%s/records/A/%s", domainName, subDomainName)
	res = d.api.HTTPExec(DELETE, url, "")

	return res
}

func (d *Domain) DeleteDNSRecord(domainName string, recordType dnsRecordType, name string) Result {
	var res Result

	rt, err := getRecordType(dnsRecordType(recordType))
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Message = err.Error()
		return res
	}

	url := fmt.Sprintf("/domains/%s/records/%v/%s", domainName, rt, name)
	res = d.api.HTTPExec(DELETE, url, "")

	return res
}

func (d *Domain) CreateDNSRecord(domainName string, recordType dnsRecordType, name string, value string, ttl int) Result {

	var res Result

	rt, err := getRecordType(dnsRecordType(recordType))
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Message = err.Error()
		return res
	}

	url := fmt.Sprintf("/domains/%s/records/%v/%s", domainName, rt, name)
	jsonStr := fmt.Sprintf(`[{"data": "%s","ttl": %d}]`, value, ttl)

	res = d.api.HTTPExec(PUT, url, jsonStr)

	return res
}

func (d *Domain) UpdateDNSRecord(domainName string, recordType dnsRecordType, name string, value string, ttl int) Result {

	var res Result

	rt, err := getRecordType(dnsRecordType(recordType))
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Message = err.Error()
		return res
	}

	url := fmt.Sprintf("/domains/%s/records/%v/%s", domainName, rt, name)
	jsonStr := fmt.Sprintf(`[{"data": "%s","ttl": %d}]`, value, ttl)

	res = d.api.HTTPExec(PUT, url, jsonStr)

	return res
}
