package main
/*
* 设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
* set(key, value)：将记录(key, value)插入该结构
* get(key)：返回key对应的value值
* [要求]
* set和get方法的时间复杂度为O(1)
* 某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
* 当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
* 若opt=1，接下来两个整数x, y，表示set(x, y)
* 若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
* 对于每个操作2，输出一个答案
 */
func main() {
	
}

func LRU( operators [][]int ,  k int ) []int {
	// write code here
	res := make([]int, 0 ,len(operators))
	key := make([]int, k)
	value := make([]int, k)
	for _,val := range operators{
		// 进入的操作我们模仿队列，先进先出
		if val[0] == 1{
			if len(key) == k{
				key = key[1:]
				value = value[1:]
			}

			key = append(key, val[1])
			value = append(value, val[2])
		} else if val[0] == 2{

			index := -1
			for i,v := range key{
				if val[1] == v{
					index = i
					break
				}
			}

			if index==-1{
				res = append(res, -1)
			}else {
				res = append(res, value[index])
				//将最近查找的这个index放到队列的尾部。表示它是最新的
				if index < k-1{

					key = append(key[:index],append(key[index+1:], key[index])...)
					value = append(value[:index],append(value[index+1:],value[index])...)
				}
			}
		}
	}

	return res
}