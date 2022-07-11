# Go Client for GoDaddy API

## Wrapper functions
```go
CreateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result
UpdateSubDomain(domainName string, subDomainName string, forwardIPAddr string) Result
DeleteSubDomain(domainName string, subDomainName string) Result
UpdateDomain(domainName string, forwardIPAddr string) Result
DomainAvailable(domainName string) Result
```

## Standard functions
```go
CreateDNSRecord(domainName string, recordType dnsRecordType, name string, value string, ttl int) Result
UpdateDNSRecord(domainName string, recordType dnsRecordType, name string, value string, ttl int) Result
DeleteDNSRecord(domainName string, recordType dnsRecordType, name string) Result
```

## How to use...
```go
package main

import (
	"fmt"

	gdc "github.com/kambahr/go-godaddyclient"
)

func main() {

    apiBaseURL := "https://api.godaddy.com"
    version:= "v1"

    apiKey := os.Getenv("GODADDY_API_KEY")
    apiSecret  := os.Getenv("GODADDY_API_SECRET")

    gdyClient := gdc.NewGoDaddyClient(apiBaseURL, version, apiKey, apiSecret)

    domainName:= "<your domain name>"
    subDomainName := "<your sub-domain name>"
    ipAddr:= "<ip address of the sub-domain will point to>"
    ttlTimeInSeconds := 600 // seconds to refresh/update the target record

    res := gdyClient.Domain.CreateDNSRecord(domainName, gdc.RecordType.A, subDomainName, ipAddr, ttlTimeInSeconds)
    fmt.Println(res)

    res = gdyClient.Domain.DomainAvailable("example.com")
    fmt.Println(res)
}
```
