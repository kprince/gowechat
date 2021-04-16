package mini

import (
	"github.com/kprince/gowechat/mp/base"
	"log"
)

const (
	aQrUrl = "https://api.weixin.qq.com/wxa/getwxacode"
	aContentUrl = "https://api.weixin.qq.com/wxa/msg_sec_check"
)

type Mini struct {
	base.MpBase
}

func (m Mini) GetAccessToken()(string,error){
	return m.Context.GetAccessToken()
}

//普通小程序码，个数限制100000个，path限制128字节，宽度：[280,1280]
func (m Mini) FilterContent(content string)([]byte,error){
	params := map[string]interface{}{
		"content":content,
	}
	response, err := m.HTTPPostJSONWithAccessToken(aContentUrl, &params)
	if err != nil {
		log.Println(err)
	}
	return response,err
}

//普通小程序码，个数限制100000个，path限制128字节，宽度：[280,1280]
func (m Mini) GetMiniCode(path string,width int, isHyaline bool)([]byte,error){
	if width <280 {
		 width = 280
	}
	if width >1280 {
		width = 1280
	}
	params := map[string]interface{}{
		"path":path,"width":width,"is_hyaline": isHyaline,
	}
	response, err := m.HTTPPostJSONWithAccessToken(aQrUrl, &params)
	if err != nil {
		log.Println(err)
	}
	return response,err
}


//小程序码，个数无限制，scene限制32字符，宽度：[280,1280]
//scene:最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
func (m Mini) GetUnLimitedMiniCode(page,scene string,width int, isHyaline bool)([]byte,error){
	if width <280 {
		 width = 280
	}
	if width >1280 {
		width = 1280
	}
	params := map[string]interface{}{
		"page":page,"scene":scene,"width":width,"is_hyaline": isHyaline,
	}
	response, err := m.HTTPPostJSONWithAccessToken(aQrUrl, &params)
	if err != nil {
		log.Println(err)
	}
	return response,err
}