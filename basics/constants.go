package basics

// whatever they say to use for constants
const pi = 3.141592653

// untyped constants
const GRAVITY = 9.81

func constants() {
	// typed constants
	const days int = 7

	// no gofer for constants
	// better to use const block like this
	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thursday  = 4
		friday    = 5
		saturday  = 6
		sunday    = 6
	)
}
