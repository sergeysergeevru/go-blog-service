package logical

//You are given a message encoded using the following mapping:
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//
//Write a function or algorithm that takes a string of digits and returns the number of ways
//it can be decoded back into its original message.
//
//For example:
//
//- Given the input "12", the possible decodings are "AB" and "L", so the output should be 2.
//- For the input "226", the possible decodings are "BZ", "VF", and "BBF", making the output 3.
//- With the input "0", there are no valid decodings, resulting in an output of 0.
//
//Your solution should efficiently handle larger inputs as well.

// GetDecodeWaysNumber Time complexity = O(len(s)) Space complexity = O(1)
func GetDecodeWaysNumber(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	prevDP := 1
	currentDP := 1
	for i := 1; i < len(s); i++ {
		digit := s[i] - '0'
		var actualDP int
		if digit != 0 {
			actualDP = currentDP
		}
		twoDigits := (s[i-1]-'0')*10 + digit
		if 10 <= twoDigits && twoDigits <= 26 {
			actualDP += prevDP
		}
		prevDP, currentDP = currentDP, actualDP
	}
	return currentDP
}
