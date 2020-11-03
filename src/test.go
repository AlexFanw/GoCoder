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
}