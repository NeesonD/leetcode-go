package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 30))
}

func coinChange(coins []int, amount int) int {
	result2 = math.MaxInt32
	dict = map[int]int{}
	back(coins, 0, amount, 0)
	if result2 == math.MaxInt32 {
		return -1
	}
	return result2
}

var result2 = math.MaxInt32

var dict = map[int]int{}

func back(coins []int, currAmount int, amount int, p int) {
	if currAmount == amount {
		result2 = int(math.Min(float64(result2), float64(p)))
		return
	}
	for _, v := range coins {
		if currAmount+v > amount {
			continue
		}
		p++
		if dict[p] == currAmount+v {
			continue
		}
		dict[p] = currAmount
		back(coins, currAmount+v, amount, p)
		p--
	}
}



