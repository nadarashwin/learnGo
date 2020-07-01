/*

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

import "fmt"

func main() {

	fmt.Println("*****------*******")
	a := []int{2, 7, 11, 15}

	// Double iteration
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			//fmt.Println(a[i], a[j])
			if a[i]+a[j] == 9 {
				fmt.Printf("the generic way will result in %d and %d\n", a[i], a[j])
			}
		}
	}

	fmt.Println("*****------*******")

	// Single iteration and then using hash for lookup
	hashTable := make(map[int]int)
	for index, element := range a {
		tempVariable := 9 - element
		if value, ok := hashTable[tempVariable]; ok {
			fmt.Printf("the combinations are %d and %d\n", a[index], a[value])
		} else {
			hashTable[element] = index
		}
	}

	d := twoSum(a, 9)
	fmt.Println(d)

}

func twoSum(nums []int, target int) []int {
	tempHash := make(map[int]int)
	for index, element := range nums {
		tempVar := target - element
		if value, ok := tempHash[tempVar]; ok {
			//fmt.Printf("%d and %d\n", nums[value], nums[index])
			return []int{value, index}
		}
		tempHash[element] = index

	}
	return []int{}
}
