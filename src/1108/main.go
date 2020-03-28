package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	ips := strings.Split(address, ".")
	return fmt.Sprintf("%s[.]%s[.]%s[.]%s", ips[0], ips[1], ips[2], ips[3])
}

func main() {
	ans := defangIPaddr("1.1.1.1")
	fmt.Println(ans)
}
