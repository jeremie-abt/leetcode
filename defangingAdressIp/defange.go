package defangingAdressIp

import "strings"

func defangIpaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}
