package tool

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
)

// reference: https://stackoverflow.com/questions/2814002/private-ip-address-identifier-in-regular-expression
const regexpPrivateIPTemplate = "(^127\\.)|(^192\\.168\\.)|(^10\\.)|(^172\\.1[6-9]\\.)|(^172\\.2[0-9]\\.)|(^172\\.3[0-1]\\.)" +
	"|(^::1$)|(^[fF][cCdD])"

var (
	regexpPrivateIP = regexp.MustCompile(regexpPrivateIPTemplate)
	privateIPBlocks []*net.IPNet
)

func init() {
	for _, cidr := range []string{
		"0.0.0.0/8",          // This host on this network       [RFC1122], section 3.2.1.3
		"10.0.0.0/8",         // Private-Use                       [RFC1918]
		"100.64.0.0/10",      // Shared Address Space              [RFC6598]
		"127.0.0.0/8",        // Loopback                          [RFC1122], section 3.2.1.3
		"169.254.0.0/16",     // Link Local                        [RFC3927]
		"172.16.0.0/12",      // Private-Use                       [RFC1918]
		"192.0.0.0/24",       // IETF Protocol Assignments         [RFC6890], section 2.1
		"192.0.0.0/29",       // IPv4 Service Continuity Prefix    [RFC7335]
		"192.0.0.8/32",       // IPv4 dummy address                [RFC7600]
		"192.0.0.9/32",       // Port Control Protocol Anycast     [RFC-ietf-pcp-anycast-08]
		"192.0.0.170/32",     // NAT64 Discovery             [RFC7050], section 2.2
		"192.0.0.171/32",     // DNS64 Discovery             [RFC7050], section 2.2
		"192.0.2.0/24",       // Documentation (TEST-NET-1)        [RFC5737]
		"192.31.196.0/24",    // AS112-v4                          [RFC7535]
		"192.52.193.0/24",    // AMT                               [RFC7450]
		"192.88.99.0/24",     // Deprecated (6to4 Relay Anycast)   [RFC7526]
		"192.168.0.0/16",     // Private-Use                       [RFC1918]
		"192.175.48.0/24",    // Direct Delegation AS112 Service   [RFC7534]
		"198.18.0.0/15",      // Benchmarking                      [RFC2544]
		"198.51.100.0/24",    // Documentation (TEST-NET-2)        [RFC5737]
		"203.0.113.0/24",     // Documentation (TEST-NET-3)        [RFC5737]
		"240.0.0.0/4",        // Reserved                          [RFC1112], section 4
		"255.255.255.255/32", // Limited Broadcast                 [RFC919], se
		"::1/128",            // IPv6 loopback
		"fe80::/10",          // IPv6 link-local
		"fc00::/7",           // IPv6 unique local addr
		"fd00::/8",           // IPv6 private ip
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

func IsPublicIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return false
		}
	}
	return true
}

func IsPrivateIP(ip string) bool {
	return regexpPrivateIP.MatchString(ip)
}

func IsIPv4(ip string) bool {
	for i := 0; i < len(ip); i++ {
		if ip[i] == '.' {
			return true
		}
	}

	return false
}

func IsIPv6(ip string) bool {
	for i := 0; i < len(ip); i++ {
		if ip[i] == ':' {
			return true
		}
	}

	return false
}

func IsValidIP(str string) bool {
	return net.ParseIP(str) != nil
}

func IsValidHostPort(str string) bool {
	host, port, err := net.SplitHostPort(str)
	if err != nil {
		return false
	}

	if !IsValidIP(host) {
		return false
	}

	num, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if num > (1 << 16) {
		return false
	}

	return true
}

func IsValidURL(str string) bool {
	_, err := url.Parse(str)
	return err == nil
}
