package main
/*
山峰元素是指其值大于或等于左右相邻值的元素。给定一个输入数组nums，任意两个相邻元素值不相等，数组可能包含多个山峰。找到索引最大的那个山峰元素并返回其索引。

假设 nums[-1] = nums[n] = -∞。
 */
func main() {
	
}
/*
这一题为什么要用到二分？？
 */
func solve(a []int) int{
	for i:=len(a)-1;i>1;i--{
		if  a[i]>=a[i-1]{
			return i
		}
	}
	return 0
}