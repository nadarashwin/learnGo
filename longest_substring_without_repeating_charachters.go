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
)

func main() {
	d := lengthOfLongestSubstring("dvdf")
	fmt.Printf("the max of the string is %d\n", d)

}

func lengthOfLongestSubstring(s string) int {
	dummyHash := make(map[string]int)
	maxLength, startPoint := 0, 0
	for index, element := range s {
		// Logic to check for duplicate
		if value, ok := dummyHash[string(element)]; ok {
			duplicatePoint := value + 1 // for a duplicate, start the index from the duplicate position + 1
			if duplicatePoint > startPoint {
				startPoint = duplicatePoint
			}
		}
		// Logic to check for max length
		tempLength := index - startPoint + 1 // index starts from 0 and length starts from 1, hence + 1
		if tempLength > maxLength {
			maxLength = tempLength
		}
		dummyHash[string(element)] = index
	}
	fmt.Printf("%v %d %d\n", dummyHash, maxLength, startPoint)
	return maxLength
}
