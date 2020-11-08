package main

func main() {
	
}
func maxWater( arr []int ) int64 {
	if len(arr) < 3 {
		return 0
	}
	left, right := 0, len(arr)-1
	leftMax, rightMax := arr[0], arr[right]
	res := 0
	for left <= right {
		if leftMax < rightMax {
			if arr[left] < leftMax {
				res += (leftMax-arr[left])
			} else {
				leftMax = arr[left]
			}
			left++
		} else {
			if arr[right] < rightMax {
				res += (rightMax-arr[right])
			} else {
				rightMax = arr[right]
			}
			right--
		}
	}
	return int64(res)
}