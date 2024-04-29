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
