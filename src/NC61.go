package main

import "fmt"

func main() {
	i := []int{0,3,2,0}
	target := 0
	fmt.Println(twoSum(i, target))
}

func twoSum( numbers []int ,  target int ) []int {
	// write code here
	var ts = make(map[int] int, len(numbers));
	for i :=range numbers{
		if j,OK := ts[target-numbers[i]];OK{
			return []int{j+1,i+1}
		}else {
			ts[numbers[i]] = i
		}
	}
	return nil;
}