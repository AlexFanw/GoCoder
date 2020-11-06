package main

import "fmt"

func main() {
	s:="abc"
	fmt.Println(permutations(s))
}

func permutations(s string)[]string{
	res := []string{}
	for key,_ := range setpermutations(s){
		res = append(res, key)
	}
	return res
}
func setpermutations(s string) map[string]int{
	if len(s) == 0{
		return map[string]int{}
	}
	res := setpermutations(s[1:])
	tmp := []string{}
	for key,_ := range res{
		tmp = append(tmp, key)
	}
	res[string(s[0])]=1 // 加入单独新元素
	for _,key := range tmp{ // 遍历已经存在的每种组合
		for i:=0;i<=len(key);i++{
			res[key[:i]+string(s[0])+key[i:]]=1
		}
	}
	return res
}