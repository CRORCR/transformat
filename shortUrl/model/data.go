package model

//长url
type Long2ShortUrlRequest struct{
	OriginUrl string `json:"origin_url"`
}

type Long2ShortUrlReponse struct{
	ShortUrl string `json:"short_url"`
	Header
}

type Short2LongRequest struct{
	ShortUrl string `json:"short_url"`
}

type Short2LongUrlReponse struct{
	OriginUrl string `json:"origin_url"`
	Header
}

type Header struct {
	Code int
	Message string
}

type ShortUrl struct{
	ShortUrl string `json:"short_url" db:"short_url"`
}

