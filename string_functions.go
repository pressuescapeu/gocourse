package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "Hello Go!"
	//// len
	//fmt.Println(len(str))
	//// concatenate
	//str2 := " Oh no I meant world"
	//result := str + str2
	//fmt.Println(result)
	//
	//// strings have index
	//// this will output the ASCII value
	//fmt.Println(str[5])
	//// slicing
	//fmt.Println(str[len(str)-3 : len(str)])
	//
	//// STANDARD library functions
	//num := 18
	//// converting to string
	//str3 := strconv.Itoa(num)
	//fmt.Println(str3 + " <- this is a number now")
	//
	//// splitting strings
	//fruits := "apple, orange, banana, pineapple"
	//parts := strings.Split(fruits, ", ")
	//fmt.Println(parts)
	//
	//// joining strings
	//countries := []string{"Germany", "USA", "Russia", "Uzbekistan"}
	//// first argument - the list
	//// second argument - the separator
	//joined := strings.Join(countries, ", ")
	//fmt.Println(joined)
	//
	//// does contain
	//// first argument - the string we are searching in
	//// second argument - the substring
	//fmt.Println(strings.Contains(str, "H"))
	//
	//// replacing
	//// we are going to get a copy of the string
	//// first argument - the string we are operating in
	//// second argument - what needs to be replaced
	//// third argument - the replacement
	//// fourth argument - how many replacements to make
	//replaced := strings.Replace(str, "Go", "World (yeah replaced it)", 1)
	//fmt.Println(replaced)
	//
	//// trimming unnecessary spaces
	//stringWithSpace := "     Hello Everyone!     "
	//fmt.Println(stringWithSpace)
	//fmt.Println(strings.TrimSpace(stringWithSpace))
	//
	//// changing the cases
	//fmt.Println(strings.ToLower(stringWithSpace))
	//fmt.Println(strings.ToUpper(stringWithSpace))
	//
	//// repeating something a fixed number of times
	//// first argument - the string that will be repeated
	//// second argument - number of repetitions
	//fmt.Println(strings.Repeat(strings.ToUpper(stringWithSpace), 4))
	//
	//// counting the substrings inside another string
	//// first argument - the string that we count in
	//// second argument - the substring
	//fmt.Println(strings.Count("Hello", "l"))
	//
	//// look for prefix / suffix
	//// first argument - the string that we search in
	//// second argument - the substring
	//fmt.Println(strings.HasPrefix("Hello", "Hell"))
	//fmt.Println(strings.HasPrefix("Hello", "Hehe"))
	//fmt.Println(strings.HasSuffix("Hello", "lo"))
	//fmt.Println(strings.HasSuffix("Hello", "ol"))
	//
	//// advanced
	//// regexp - screw THAT bro
	//str5 := "Hello, 123 Go! 11"
	//// a function that compiles a regexp, returns a regexp
	//// defining the pattern, d means digits, + means more than 1
	//re := regexp.MustCompile(`\d+`)
	//// find all strings that match the pattern in re
	//// arguments are the string and number of matches that we want
	//// if -1 that will give all the matches
	//matches := re.FindAllString(str5, -1)
	//fmt.Println(matches)
	//
	//// utf8 package
	//str6 := "Че там"
	//// counting runes
	//fmt.Println(utf8.RuneCountInString(str6))

	// string builder is a type in Go
	// efficient concatenating strings in terms of memory

	var builder strings.Builder

	// some sample strings
	builder.WriteString("So we start here")
	builder.WriteString(", continue here")
	builder.WriteString(", end here")

	// the final string - convert
	result := builder.String()

	fmt.Println(result)

	// WriteRune - add characters sequentially
	builder.WriteRune(' ')
	builder.WriteString("and boom we got here")

	result = builder.String()

	fmt.Println(result)

	// start a new string? reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")

	result = builder.String()
	fmt.Println(result)
}
