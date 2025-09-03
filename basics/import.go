package basics

import (
	"fmt"
	foo "net/http"
)

func import_learn() {
	fmt.Println("Hello World!")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status:", resp.Status)
}
