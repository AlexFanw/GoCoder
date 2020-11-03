package main

import "fmt"

/*
请实现有重复数字的有序数组的二分查找。
输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一。
5,4,[1,2,4,4,5]
 */
func main() {
	var ls = []int{2,3,4,4,4,7,7,8,10,10,11,12,13,14,15,15,17,18,19,23,24,24,24,24,25,26,26,26,27,27,28,29,29,30,33,36,38,38,40,40,41,43,43,43,44,46,46,47,51,52,52,53,54,56,57,57,57,58,58,61,61,61,62,64,64,66,66,67,67,67,70,72,74,74,74,75,75,78,78,78,79,79,80,83,83,83,83,84,84,86,88,89,89,90,91,91,92,93,93,96};
	fmt.Println(upper_bound_(len(ls),1, ls));
}

/*
这题要注意，不是找数组里第一个跟它一样大的数，是第一个大于等于它的数！
 */
func upper_bound_( n int ,  v int ,  a []int ) int {
	// write code here
	if len(a)==0 {
		return 1;
	}
	var index int = 0;
	for ;;{
		var halflen int = len(a)/2;
		if a[0]>=v &&len(a)==1{
			return index+1;
		}
		if a[halflen-1] >= v{
			a = a[0:halflen];
			//fmt.Print("a",index)
			//fmt.Println(a)
		} else if a[len(a)-1] >= v && a[halflen] <= v{
			index = index+halflen;
			a = a[halflen:len(a)];
			//fmt.Print("b",index)
			//fmt.Println(a)
		} else{
			return n+1
		};
	}
}
