package model

//长地址-->短地址 请求
type Long2ShortRequest struct {
	OriginUrl string `json:"origin_url"`
}

//长地址-->短地址 回报
type Long2ShortResponse struct {
	Header
	ShortUrl string `json:"short_url"`
}

type Header struct{
	Code int `json:"code"`
	Message string `json:"message"`
}

//短地址转长地址
//短地址-->长地址 请求
type Short2LongRequest struct {
	ShortUrl string `json:"short_url"`
}

//短地址-->长地址 回报
type Short2LongResponse struct {
	Header
	OriginUrl string `json:"origin_url"`
}