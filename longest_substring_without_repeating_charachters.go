/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	d := lengthOfLongestSubstring("abccbde")
	fmt.Printf("the max of the string is %d\n", d)

}

func lengthOfLongestSubstring(s string) int {
	dummYHash := make(map[string]int)
	var dummyString string
	for _, element := range s {
		if strings.Contains(dummyString, string(element)) {
			dummYHash[dummyString] = len(dummyString)
			dummyString = string(element)
		} else {
			dummyString = dummyString + string(element)
		}
		dummYHash[dummyString] = len(dummyString)

	}
	fmt.Printf("%s and %v\n", dummyString, dummYHash)

	return getMax(dummYHash)
}

func getMax(a map[string]int) int {
	var max int
	for _, value := range a {
		if max < value {
			max = value
		}
	}
	return max
}
