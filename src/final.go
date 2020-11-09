package main

import "fmt"

func main() {
	arr := []int{5,3,1,4,2,3}
	fmt.Println(solution(arr))
}

func solution(arr []int)[]int{
	if len(arr) == 0{
		return []int{}
	}
	res := make([]int,len(arr))
	stack := []int{}
	for ind,val := range arr{
		if ind==0{
			stack = append(stack,ind)
			continue
		}
		if val <= arr[stack[len(stack)-1]]{
			stack = append(stack, ind)
		}
		for val > arr[stack[len(stack)-1]]{
			if val > stack[len(stack)-1] {
				res[stack[len(stack)-1]] = val
				stack = stack[:len(stack)-1] //出栈
			}
		}
		stack = append(stack, ind)
	}
	return res
}
