package myutils

import (
	"fmt"
	"net"
	"testing"
)

func TestLookup(t *testing.T) {
	dns := "tans.fun"
	ip, err := net.LookupIP(dns)

	if err != nil {
		fmt.Print("reslove error")
	}
	fmt.Println(ip)
}

func TestParseDomainName(t *testing.T) {
	lookup, err := DnsLookup("tans.fun", "192.5.6.30")
	if err != nil {
		return
	}
	fmt.Println(lookup)

}

func TestSendUpdate(t *testing.T) {
	type args struct {
		domain    string
		ip        string
		dnsServer string
		dnsPort   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "test insert", args{domain: "google.com", ip: "1.1.1.1", dnsServer: "127.0.0.1", dnsPort: "30053"}},
		{n},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendUpdate(tt.args.domain, tt.args.ip, tt.args.dnsServer, tt.args.dnsPort)
		})
	}
}
