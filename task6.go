package main

import "fmt"

func SimpleEquations(a, b, c int) {

	// your code here
	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000-x; y++ {
			for z := 1; z <= 10000/(x*y); z++ {
				if x+y+z == a {
					if x*y*z == b {
						if x*x+y*y+z*z == c {
							fmt.Println(x, y, z)
							return
						}
					}
				}
			}
		}
	}

	fmt.Println("no solution")
	return

}

func moneyCoins(money int) []int {

	// your code here
	Uang := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	kembalian := []int{}
	// your code here
	for i := len(Uang) - 1; i >= 0; i-- {
		for Uang[i] <= money {
			kembalian = append(kembalian, Uang[i])
			money -= Uang[i]
			if money == 0 {
				return kembalian
			}
		}
	}
	return kembalian

}

func delete(slice []int, s int) {
	slice = append(slice[:s], slice[s+1:]...)
}
func DragonOfLoowater(dragonHead, knightHeight []int) {

	// your code here
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
		return
	}

	var minTotHeight, tempH int
	var isKilled bool
	var index int

	for _, diameter := range dragonHead {
		tempH = 0
		isKilled = false
		for i, height := range knightHeight {
			if height >= diameter && height < tempH {
				minTotHeight += height - tempH
				tempH = height
				index = i
				continue
			} else if isKilled {
				continue
			} else if height >= diameter {
				minTotHeight += height
				tempH = height
				index = i
				isKilled = true
			}
		}

		if !isKilled {
			fmt.Println("knight fall")
			return
		}

		delete(knightHeight, index)
	}

	fmt.Println(minTotHeight)
	return

}

func BinarySearch(array []int, x int) {

	// your code her
	left, right := 0, len(array)-1

	if x > array[len(array)-1] || x < array[0] {
		fmt.Println(-1)
		return
	} else if array[left] == x {
		fmt.Println(left)
		return
	} else if array[right] == x {
		fmt.Println(right)
		return
	} else {
		for left+1 != right {
			if array[(left+right)/2] == x {
				fmt.Println((left + right) / 2)
				return
			} else if array[(left+right)/2] < x {
				left = (left + right) / 2
			} else {
				right = (left + right) / 2
			}
		}
	}

	fmt.Println(-1)
	return

}

func main() {
	// Simple Equations
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3

	// Money Coins
	fmt.Println(moneyCoins(123))   // [100 20 1 1 1]
	fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]
	fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]
	fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

	// Dragon of Loowater
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})      // 11
	DragonOfLoowater([]int{5, 10}, []int{5})           // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})   // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})   // 10
	DragonOfLoowater([]int{10, 9}, []int{10, 1, 8, 5}) // knight fall

	// Binary Search
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  // 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 // 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  // 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
