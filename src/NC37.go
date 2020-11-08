package main

func main() {
	
}
/*
1.以区间左边界为基准排序区间
2.滑动窗口，start窗口左边界，end窗口右边界
3.往前滑动，当end不在下一个区间之中时，得到一个合并区间停止滑动，移动到下一个区间接着滑动
 */
func merge37( intervals []*Interval ) []*Interval {
	// write code here
	for i:=1; i< len(intervals); i++ {
		for j:=i;j>0;j-- {
			if intervals[j].Start < intervals[j-1].Start {
				intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
			}
		}
	}

	for i:=1; i< len(intervals); i++ {
		if intervals[i-1].End < intervals[i].Start {
			//两个区间保持不变
		}else if intervals[i-1].End >= intervals[i].Start && intervals[i-1].End <= intervals[i].End{
			//两个区间合并
			intervals[i].Start = intervals[i-1].Start
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		}else if intervals[i-1].End > intervals[i].End {
			intervals[i].Start = intervals[i-1].Start
			intervals[i].End = intervals[i-1].End
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		}
	}

	return intervals
}