package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", printFoo)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}


func printFoo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello world")
}