package main

func main() {
	
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, num := range nums {
		index,ok := dict[target-num]
		if ok {
			return []int{i,index}
		}
		dict[num] = i
	}
	return nil
}
