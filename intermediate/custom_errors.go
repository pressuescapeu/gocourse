package intermediate

import (
	"errors"
	"fmt"
)

type customError struct {
	code       int
	message    string
	err_custom error
}

// Error returns the error message, implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.err_custom)
}

// doSomething - function that returns a custom error
func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:       500,
			message:    "Something went wrong",
			err_custom: err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}

func custom_errors() {
	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}
	// won't happen
	fmt.Println("All good")
}
