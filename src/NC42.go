package main
import "sort"
func main() {
	
}
func permuteUnique( num []int ) [][]int {
	rs :=[][]int{}
	var backtrace func([]int,[]bool)
	backtrace = func(path []int, vis []bool){
		if len(path)==len(num){
			cp := make([]int,len(num))
			copy(cp,path)
			rs = append(rs,cp)
			return
		}
		for i:=0;i<len(num);i++ {
			if vis[i] || (i!=0 && num[i-1]==num[i] && !vis[i-1]){
				continue
			}
			vis[i]=true
			backtrace(append(path,num[i]), vis)
			vis[i]=false
		}
	}
	sort.Ints(num)
	backtrace([]int{},make([]bool,len(num)))
	return rs
}