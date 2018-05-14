package engine

type Request struct {
	Url string
	//第二个参数是函数 传入byte 传出result
	ParseFunc func([]byte) ParseResult
}

//返回结果是多个请求和多个item
type ParseResult struct {
	Request []Request
	Item    []interface{}
}

func NewParse([]byte) ParseResult {
	return ParseResult{}
}
