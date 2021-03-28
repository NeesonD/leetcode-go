package main

import (
	"fmt"
	"math"
)

func main() {
 fmt.Println(minEatingSpeed([]int{3,6,7,11},8))
}


func shipWithinDays(weights []int, D int) int {
	var left int
	var right int
	var mid int
	for _, weight := range weights {
		left = int(math.Max(float64(left), float64(weight)))
		right += weight
	}
	for left < right {
		mid = (left + right) >> 1
		if canShip(weights, D, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canShip(weights []int, D int, k int) bool {
	cur := k
	for _, weight := range weights {
		if weight > k {
			return false
		}
		if weight > cur {
			cur = k
			D--
		}
		cur -= weight
	}
	return D > 0
}



func minEatingSpeed(piles []int, H int) int {
	left := 1
	var right int
	for _, pile := range piles {
		right = int(math.Max(float64(right), float64(pile)))
	}
	for left < right{
		mid := (left+right) >> 1
		if canEat(piles, H, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canEat(piles []int, H int,k int) bool  {
	var total int
	for _, pile := range piles {
		if pile%k > 0{
			total += pile/k + 1
		} else {
			total += pile/k
		}
	}
	return total <= H
}
