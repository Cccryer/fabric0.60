package myutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/miekg/dns"
	"net"
	"strings"
)

type DNSHeader struct {
	ID            uint16 // 标识
	Flag          uint16
	QuestionCount uint16
	AnswerRRs     uint16 //RRs is Resource Records
	AuthorityRRs  uint16
	AdditionalRRs uint16
}

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

func reslove(name string) (addrs []net.IP, err error) {
	return nil, nil
}

func main() {
	addr := net.UDPAddr{
		Port: 53,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Failed to set up server:", err)
		return
	}

	for {
		handleRequest(conn)
	}
}

func handleRequest(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	_, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Failed to read from UDP:", err)
		return
	}

	// Parse the DNS query here, and get the domain name.
	domain := parseQuery(buffer)
	// Send a query to the root server.
	response := queryRootServer(domain)
	// Write the response back to the client.
	conn.WriteToUDP(response, addr)
}

func parseQuery(buffer []byte) string {
	// This is a placeholder. In a real DNS server, you would parse the DNS query
	// to get the domain name.
	return strings.TrimRight(string(buffer), "\x00")
}
func queryRootServer(domain string) []byte {
	// This is a placeholder. In a real DNS server, you would send a DNS query
	// to the root server and return the response.
	return []byte("This is a placeholder response for domain: " + domain)
}

//要将域名解析成相应的格式，例如：
//"www.google.com"会被解析成"0x03www0x06google0x03com0x00"
//就是长度+内容，长度+内容……最后以0x00结尾
func ParseDomainName(domain string) []byte {
	var (
		buffer   bytes.Buffer
		segments []string = strings.Split(domain, ".")
	)
	for _, seg := range segments {
		binary.Write(&buffer, binary.BigEndian, byte(len(seg)))
		binary.Write(&buffer, binary.BigEndian, []byte(seg))
	}
	binary.Write(&buffer, binary.BigEndian, byte(0x00))

	return buffer.Bytes()
}

func buildDnsPackete(domain string) {
	pre := []byte{0x01, 0x80, 1, 0, 0, 0}
	tail := []byte{1, 1}
	var payload = []byte{}
	for _, s := range strings.Split(domain, ".") {
		payload = append(payload, byte(len(s)))
		payload = append(payload, []byte(s)...)
	}
	payload = append(payload, byte(0x00))

	query := append(pre, payload...)
	query = append(query, tail...)

}
