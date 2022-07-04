package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//负责业务逻辑，遇到错误只返回，不处理错误
func HttpFileListHelper(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/filelist/"):]
	fmt.Println(r.URL)
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(content)
	return nil
}
