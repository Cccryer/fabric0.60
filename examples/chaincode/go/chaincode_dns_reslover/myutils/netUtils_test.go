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
	lookup, err := DnsLookup("example.com", "127.0.0.1", "30054")
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
		{
			name: "test insert",
			args: args{
				domain:    "abc.google.com",
				ip:        "1.1.1.2",
				dnsServer: "127.0.0.1",
				dnsPort:   "30054",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendUpdate(tt.args.domain, tt.args.ip, tt.args.dnsServer, tt.args.dnsPort)
		})
	}
}
