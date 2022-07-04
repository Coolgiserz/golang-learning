package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	"coolgiserz.com/learngo/errhandling/filelistingserver/filelisting"
	"github.com/gpmgo/gopm/log"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

//统一处理错误，再返回一个能够作为http.HandleFunc参数的返回值。errWrapper将函数作为参数，将函数作为返回值
func errHandler(app appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := app(w, r)
		if err != nil {
			log.Warn("Error request %s\n", err) //gopm模块下的log功能
			// log.Printf("Error request %s\n", err) //标准库的日志模块log
			code := http.StatusOK
			switch {
			case os.IsNotExist(err): //无文件

				code = http.StatusNotFound
			case os.IsPermission(err): //没权限
				code = http.StatusForbidden
			default: //内部错误
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

//能否进一步封装，把错误处理部分与正常业务逻辑分离？
func main() {
	fmt.Println("Web Server")
	// http.HandleFunc("/filelist/",
	// 	func(w http.ResponseWriter, r *http.Request) { //func这个匿名函数可以封装成指处理业务逻辑、处理错误细节的方法
	// 		path := r.URL.Path[len("/filelist/"):]
	// 		fmt.Println(r.URL)
	// 		file, err := os.Open(path)

	// 		if err != nil {
	// 			// panic(err)
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 			return
	// 		}
	// 		defer file.Close()
	// 		content, err := ioutil.ReadAll(file)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		w.Write(content)

	// 	})
	http.HandleFunc("/filelist/", errHandler(filelisting.HttpFileListHelper))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
