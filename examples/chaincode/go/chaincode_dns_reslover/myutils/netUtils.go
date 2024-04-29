package myutils

import (
	"fmt"
	"github.com/miekg/dns"
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
