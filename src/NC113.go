package main

import (
	"strconv"
	"strings"
)

func main() {
	
}

func solve113(IP string) string {
	// write code here
	if len(IP) <= 15 {
		//IPV4
		ipArr := strings.Split(IP, ".")
		if len(ipArr) == 4 {
			for i := 0; i < 4; i++ {
				if v, err := strconv.Atoi(ipArr[i]); err == nil && v < 256 && ipArr[i][0] != '0' {
					continue
				}
				return "Neither"
			}
			return "IPv4"
		}
	} else {
		//IPV6
		ipArr := strings.Split(IP, ":")
		if len(ipArr) == 8 {
			for i := 0; i < 8; i++ {
				if len(ipArr[i]) <= 4 {
					for j := 0; j < len(ipArr[i]); j++ {
						if (ipArr[i][j] >= '0' && ipArr[i][j] <= '9') ||
							(ipArr[i][j] >= 'a' && ipArr[i][j] <= 'z') ||
							(ipArr[i][j] >= 'A' && ipArr[i][j] <= 'Z') {
							continue
						}
						return "Neither"
					}
					continue
				}
				return "Neither"
			}
			return "IPv6"
		}
	}
	return "Neither"
}
