package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errNotFound(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrNotExist
}

//httptest 专门用于测试http服务的标准库
/*
 构造测试用例，每个测试用例输入appHandler,检测返回的状态码code和消息message
*/
func TestErrHandler(t *testing.T) {
	tests := []struct {
		a       appHandler
		code    int
		message string
	}{
		{errNotFound, 404, "Not Found"},
	}

	for _, test := range tests {
		toTestFunc := errHandler(test.a) //返回一个函数
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		toTestFunc(response, request)
		verifyResponse(response.Result(), test.code, test.message, t)
	}

}

//验证http响应中的状态码和消息是否与预期一致
func verifyResponse(r *http.Response, code int, message string, t *testing.T) {
	b, _ := ioutil.ReadAll(r.Body)
	body := string(b)
	body = strings.Trim(body, "\n")
	if r.StatusCode != code || string(body) != message {
		t.Errorf("\nError, expected code: %d, message: %s \n got code %d, message %s", code, message, r.StatusCode, body)
	}
}
