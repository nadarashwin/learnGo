/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
package main

import (
	"fmt"
)

func main() {
	x := 1534236469
	// getExpo := expo(2, 31)
	// if x <= -getExpo || x >= getExpo-1 {
	// 	fmt.Println("champa")
	// }
	// fmt.Println("the expo value is ", getExpo)
	if x-int(int32(x)) > 0 { //overflow
		fmt.Println("Overflow Mode")
	}
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	initialSet := 0
	for x > 0 {
		reminder := x % 10
		initialSet = initialSet*10 + reminder
		x = x / 10
	}
	if initialSet-int(int32(initialSet)) > 0 { //overflow
		fmt.Println("Overflow Mode")
	}
	// if initialSet <= -getExpo || initialSet >= getExpo-1 {
	// 	fmt.Println("champa")
	// }
	if negative {
		fmt.Println(-(initialSet))
	} else {
		fmt.Println(initialSet)
	}
}

// Could have used math.Pow to get the exponential but the return value of math.Pow is float.
// Used expo to get the max and min value of 32 bit integer, however figured out a way by using
// a if condition.
func expo(base, expo int) int {
	if expo < 0 {
		return expo
	}
	tempAssig := 1
	for expo > 0 {
		tempAssig *= base
		expo--
	}
	return tempAssig
}

// Python Flow

// def reverse(x):
// 	if x <= -2**31 or x >= 2**31-1:
// 		return 0
// 	champa = False
// 	if x < 0:
// 		champa = True
// 		x = -x
// 	print(x)
// 	duplicateSet = 0
// 	while x > 0:
// 		reminder = x % 10
// 		duplicateSet = duplicateSet * 10 + reminder
// 		x = x // 10
// 	print(duplicateSet)
// 	if duplicateSet <= -2**31 or duplicateSet >= 2**31-1:
// 		return 0
// 	if champa:
// 		return -(duplicateSet)
// 	else:
// 		return duplicateSet
