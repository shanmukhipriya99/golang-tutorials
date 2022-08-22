package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://gogle.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} 
	// else {
	// 	fmt.Println(resp) //resp is a struct
	// }

	// bs := make([]byte, 99999) //not the right practice always
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
	
}