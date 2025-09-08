package basics

import (
	"fmt"
	"maps"
)

func maps_main() {
	myMap := make(map[string]int)
	fmt.Println(myMap) // empty, just map[]
	myMap["key1"] = 90 // double quotes!!!
	myMap["key2"] = 111
	fmt.Println(myMap)

	// delete key value pair
	delete(myMap, "key2")
	fmt.Println(myMap)

	// clear the whole map
	myMap["key1"] = 90
	myMap["key2"] = 111
	myMap["key3"] = 333
	fmt.Println(myMap)
	clear(myMap)
	fmt.Println(myMap)

	myMap["key1"] = 90
	myMap["key2"] = 111
	myMap["key3"] = 333
	// so wtf this unknownValue ?
	// this is a boolean which essentially tells you that value indeed exists in the map
	value, unknownValue := myMap["key1"]
	fmt.Println(value, unknownValue)
	// convention is to use ok as variable name
	value1, ok := myMap["key4"]
	fmt.Println(value1, ok) // so this is 0 false
	// because there is no key4

	myMap2 := map[string]int{"a": 1, "b": 2}
	// checking map equality
	if maps.Equal(myMap, myMap2) {
		fmt.Println("equal")
	} else {
		fmt.Println("nah")
	}

	for k, v := range myMap2 {
		fmt.Println(k, v)
	}
	// empty map
	var myMap3 map[string]string
	fmt.Println(myMap3)
	if myMap3 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("nah")
	}
	// so turns out can't do this
	// why? bc when we declare a map with var, it just points to a nil
	// so myMap3 points to a nil
	// there is no hashMap backing this
	//myMap3["key"] = "value"
	//fmt.Println(myMap3)

	// but here, when you initialize with make
	myMap4 := make(map[string]string)
	// make actually allocates and sets up the internal hash table structure
	// so now it's real and usable
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	// nested map
	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
