package reverse

// Source: exercism/problem-specifications
// Commit: 6c95c2e reverse-string: Add a test with an even-sized word (#1519)
// Problem Specifications Version: 1.2.0

type reverseTestCase struct {
	description string
	input       string
	expected    string
}

var testCases = []reverseTestCase{
	{
		description: "an empty string",
		input:       "",
		expected:    "",
	},
	{
		description: "a word",
		input:       "robot",
		expected:    "tobor",
	},
	{
		description: "a capitalized word",
		input:       "Ramen",
		expected:    "nemaR",
	},
	{
		description: "a sentence with punctuation",
		input:       "I'm hungry!",
		expected:    "!yrgnuh m'I",
	},
	{
		description: "a palindrome",
		input:       "racecar",
		expected:    "racecar",
	},
	{
		description: "an even-sized word",
		input:       "drawer",
		expected:    "reward",
	},
	// Added test case with non-ASCII unicode runes
	{
		description: "a string with non-ASCII unicode runes",
		input:       "Hello, 世界",
		expected:    "界世 ,olleH",
	},
	{
		description: "accent unicode",
		input:       "bròwn",
		expected:    "nwòrb",
	},
}
