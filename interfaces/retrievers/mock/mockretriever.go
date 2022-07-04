package mock

type Retriver struct {
	Content string
}

//实现者不需要声明自己实现了哪个接口，只需要实现接口中的方法（方法名、参数、返回值一致）即可
func (r Retriver) Get(url string) string {
	return r.Content
}

func (r Retriver) Test() string {
	// io.Reader
	// os.Open()
	// fmt.Stringer
	return "Hello interface"
}
