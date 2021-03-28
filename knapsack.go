package main

import "math"

func main() {

}

var k int

var knapsackResult = 0

func knapsack(org []int) int {
	var selectItem []int
	backtrack(selectItem, org)
	return knapsackResult
}

func backtrack(selectItem []int, canSelectItem []int) {
	sum := getSum(selectItem)
	for index, i := range canSelectItem {
		if i == 0 {
			continue
		}
		if sum+i > k {
			knapsackResult = int(math.Max(float64(knapsackResult), float64(sum)))
			continue
		}
		selectItem = append(selectItem, i)
		canSelectItem[index] = 0
		backtrack(selectItem, canSelectItem)
		selectItem = selectItem[:len(selectItem)-1]
	}
}

func getSum(selectItem []int) int {
	var sum int
	for _, i := range selectItem {
		sum += i
	}
	return sum
}
