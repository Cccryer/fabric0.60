package myutils

import (
	"fmt"
	"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_dns_reslover/miekg/dns"
)

// DnsLookup 根据指定的DNS服务器查询相应的IP地址
func DnsLookup(domain string, dnsServer string) ([]string, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	m.RecursionDesired = true

	r, _, err := c.Exchange(m, dnsServer+":53")

	if err != nil {
		return nil, err
	}

	if r.Rcode != dns.RcodeSuccess {
		return nil, fmt.Errorf("dns query failed")
	}

	var ips []string
	for _, a := range r.Answer {
		if t, ok := a.(*dns.A); ok {
			ips = append(ips, t.A.String())
		}
	}
	return ips, nil
}

// SendUpdate 构造 DNS 更新命令，更新记录
func SendUpdate(domain string, ip string, dnsServer string, dnsPort string) {
	var m dns.Msg
	m.SetUpdate(domain)

	var newRRs []dns.RR
	record := fmt.Sprintf("%s 60 IN A %s", domain, ip)
	rr, _ := dns.NewRR(record)
	newRRs = append(newRRs, rr)
	m.Insert(newRRs)
	var client dns.Client
	res1, res2, err := client.Exchange(&m, fmt.Sprintf("%s:%s", dnsServer, dnsPort))
	fmt.Println(res1, res2, err)
}
