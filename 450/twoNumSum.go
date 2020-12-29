package main


func main(){

}

func sum(arr []int, target) []int {
	nums:= map[int]bool{}

	for _, num:= range arr {
		y := target - num

		if _, found := nums[y]; found {
			return []int{y, num}
		}
		nums[y] = true
	}
	return []int{}
}