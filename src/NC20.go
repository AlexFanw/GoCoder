package main

import "strconv"

func main() {
	
}
func restoreIpAddresses(s string) []string {
	// write code here
	return split(s, 4)
}

func split(s string, n int) (ret []string) {
	l := len(s)
	if n < 1 || l < n || n > 4 {
		return nil
	}
	if n == 1 {
		if validate(s) {
			ret = append(ret, s)
		}
		return
	}
	if s[0] == '0' {
		res := split(s[1:], n-1)
		if res == nil {
			return
		}
		for _, v := range res {
			ret = append(ret, "0."+v)
		}
		return
	}
	// s0 == '1'-'9'
	for i := 1; i < l && i <= 3; i++ {
		if validate(s[0:i]) {
			res := split(s[i:], n-1)
			if res != nil {
				for _, v := range res {
					ret = append(ret, s[0:i]+"."+v)
				}
			}
		}
	}

	return
}

func validate(s string) bool {
	l := len(s)
	if l == 0 || l > 3 {
		return false
	}
	if l > 1 && s[0] == '0' {
		return false
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if i >= 0 && i <= 255 {
		return true
	}
	return false
}
