package intermediate

import (
	"errors"
	"fmt"
)

// defining our custom errors

type myError struct {
	message string
}

// here, Error is with capital E because Error is an interface
// Go has an error interface with a method of Error

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprosses() error {
	return &myError{"Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("read data: %w", err)
	}

	return nil
}

func readConfig() error {
	return errors.New("config error")
}

//// sqrt takes in float64, outputs float64 and error in case of an error
//func sqrt(x float64) (float64, error) {
//	if x < 0 {
//		return 0, errors.New("Math: square root of negative number")
//	}
//	return 1, nil
//}
//
//func process(data []byte) error {
//	if len(data) == 0 {
//		return errors.New("Error: empty data")
//	}
//	return nil
//}

func errors_not_custom() {
	//result, err := sqrt(16)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(result)
	//
	//result1, err1 := sqrt(-22)
	//if err1 != nil {
	//	fmt.Println(err1)
	//	return
	//}
	//fmt.Println(result1)

	//er := eprosses()
	//if er != nil {
	//	fmt.Println("Error: ", er)
	//	return
	//}
	//fmt.Println("Data processed successfully")
	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data read successfully")

	// panic only when error is unrecoverable
	// mostly use error
	// log errors in a proper log file
}
