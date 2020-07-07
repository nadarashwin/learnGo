/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
package main

import "fmt"

// Expand around center approach
func main() {
	var longestPalindromicString string
	testString := "racecar"
	for i := range testString {
		odd := expandFromCenter(testString, i, i)
		even := expandFromCenter(testString, i, i+1)
		fmt.Printf("the value of odd is %s and value of even is %s\n", odd, even)

		tempHolder := ""
		if len(odd) > len(even) {
			tempHolder = odd
		} else {
			tempHolder = even
		}

		if len(tempHolder) > len(longestPalindromicString) {
			longestPalindromicString = tempHolder
		}

	}
	fmt.Printf("The longest palindromic substring is %s\n", longestPalindromicString)

}

func expandFromCenter(s string, start, end int) string {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return s[start+1 : end]
}

/* Output of the above snip
╭─~/golang/code using ‹ruby-2.1.9›
╰─○ go run src/github.com/nadarashwin/learnGo/longest_palindromic_substring.go
the value of odd is a and value of even is aa
the value of odd is a and value of even is
the value of odd is b and value of even is aabbaa
the value of odd is b and value of even is
the value of odd is a and value of even is aa
the value of odd is a and value of even is
The longest palindromic substring is aabbaa
╰─○ go run src/github.com/nadarashwin/learnGo/longest_palindromic_substring.go
the value of odd is r and value of even is
the value of odd is a and value of even is
the value of odd is c and value of even is
the value of odd is racecar and value of even is
the value of odd is c and value of even is
the value of odd is a and value of even is
the value of odd is r and value of even is
The longest palindromic substring is racecar
*/

/* This is the brute force way of doing in python
class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        m = ''
        for i in range(len(s)):
            for j in range(len(s), i, -1):
                if len(m) >= j-i: # Max length is already found so no need to iterate through the shorter ones
                    break
                if s[i:j] == s[i:j][::-1]:
                    m = s[i:j]
		return m
*/
