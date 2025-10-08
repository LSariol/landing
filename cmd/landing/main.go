package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	fmt.Println("Running on port 3000")
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		panic(err)
	}

}
