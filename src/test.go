package main
import "fmt"

func main() {
	month := 2
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	switch month {
	case 3, 4, 5:
	fmt.Println("春天")
	case 6, 7, 8:
	fmt.Println("夏天")
	case 9, 10, 11:
	fmt.Println("秋天")
	case 12, 1, 2:
	fmt.Println("冬天")
	default:
	fmt.Println("输入有误...")
	}
	t1 := []int{1}
	t1 = make([]int, 0)
	fmt.Println(t1)
	a := []int{2,51,12,95,42,52,76,77,23,81,71,41,2,23,43,4,64,22,71,96,1,87,51,91,67,16,58,11,44,38,63,14,4,69,88,49,92,91,9,15,17,74,21,91,24,78,62,50,82,26,53,18,25,14,94,79,44,11,36,38,44,53,9,34,58,6,50,82,81,50,36,1,6,61,9,47,33,47,84,41,57,48,73,18}
	fmt.Println(a[82])
	fmt.Println(3<<1)
	ss := "alexfan"
	fmt.Println(ss[0:1])
}