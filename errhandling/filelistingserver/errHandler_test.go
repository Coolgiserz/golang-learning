package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

//测试用自定义错误类型,实现userError接口
type testCustomedError string

func (e testCustomedError) Error() string {
	return e.Message()
}

func (e testCustomedError) Message() string {
	return string(e)
}
func errNotFound(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrNotExist
}

func errIsPermission(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrPermission
}

func errInternalServerError(_ http.ResponseWriter, _ *http.Request) error {
	return errors.New("Internal error")
}

func errCustomedError(_ http.ResponseWriter, _ *http.Request) error {
	return testCustomedError("User error")
}

//全局定义的公共测试用例，可被TestErrHandler和TestErrHandlerInServer的测试流程共用
var tests = []struct {
	a       appHandler
	code    int
	message string
}{
	{errNotFound, 404, "Not Found"},
	{errIsPermission, 403, "Forbidden"},
	{errInternalServerError, 500, "Internal"},
	{errCustomedError, 400, "User error"},
}

//httptest 专门用于测试http服务的标准库
/*
	测试errHandler函数

	通过构造的http请求测试，通过函数接口测试
 构造测试用例，每个测试用例输入appHandler,检测返回的状态码code和消息message
*/
func TestErrHandler(t *testing.T) {

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
		t.Errorf("\nError, expected (code, message): (%d, %s) \n got (code, message): (%d, %s)", code, message, r.StatusCode, body)
	}
}

//还可以在服务器中对http服务进行测试
func TestErrHandlerInServer(t *testing.T) {
	for _, test := range tests {
		f := errHandler(test.a)
		server := httptest.NewServer(http.HandlerFunc(f)) //http.HandlerFunc将f转换为Server的一个接口类型，Handler接口定义见server.go 86行
		resp, _ := http.Get(server.URL)                   //向server服务器发送http请求
		verifyResponse(resp, test.code, test.message, t)
	}
}
