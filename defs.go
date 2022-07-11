package godaddyclient

type Result struct {
	StatusCode int
	Message    string
}

type GoaddyClient struct {
	api IGodaddyAPI

	Domain IDomain
}

type GodaddyAPI struct {
	BaseURL string
	Version string
	Key     string
	Secret  string
}

var RecordType DNSRecordType

type DNSRecordType struct {
	A     dnsRecordType
	AAAA  dnsRecordType
	CName dnsRecordType
	MX    dnsRecordType
	TXT   dnsRecordType
	SRV   dnsRecordType
	CAA   dnsRecordType
	NS    dnsRecordType
}

type dnsRecordType byte

const (
	a     dnsRecordType = 0
	aaaa  dnsRecordType = 1
	cname dnsRecordType = 2
	mx    dnsRecordType = 3
	txt   dnsRecordType = 4
	srv   dnsRecordType = 5
	caa   dnsRecordType = 6
	ns    dnsRecordType = 7
)

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)
