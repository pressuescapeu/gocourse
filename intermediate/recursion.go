// we generally use recursion when we traverse through trees!!
// and the rest that I don't really need to know yk
// memoization - remembering results (not built-in)
// keep a map of already computed results
// so if the function sees an input it already has solved, it just returns instead of overcalculating
package intermediate

var memo = make(map[int]int)

func fibMemo(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	result := fibMemo(n-1) + fibMemo(n-2)
	memo[n] = result
	return result
}

func factorial(n int) int {
	// base case: factorial of 0 is 1
	if (n == 0) || (n == 1) {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}
