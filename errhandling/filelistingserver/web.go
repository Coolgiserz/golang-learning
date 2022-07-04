package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Web Server")
	http.HandleFunc("/filelist/",
		func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path[len("/filelist/"):]
			fmt.Println(r.URL)
			file, err := os.Open(path)

			if err != nil {
				panic(err)
			}
			defer file.Close()
			content, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			w.Write(content)

		})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
