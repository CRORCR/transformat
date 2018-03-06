package logic

import (
	"transformat/short_url/model"
)

func Long2Short(req *model.Long2ShortRequest)(r *model.Long2ShortResponse,err error){
	reponse:=&model.Long2ShortResponse{}
	reponse.ShortUrl=req.OriginUrl
	return

}
