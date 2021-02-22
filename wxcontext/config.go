package wxcontext

import "github.com/astaxie/beego/cache"

// Config for user
type Config struct {
	AppID          string
	AppSecret      string
	MiniAppID          string //小程序appid
	MiniAppSecret      string //小程序secret
	Token          string
	EncodingAESKey string
	Cache          cache.Cache

	//mch商户平台需要的变量
	//证书
	SslCertFilePath string //证书文件的路径
	SslKeyFilePath  string
	SslCertContent  string //证书的内容
	SslKeyContent   string
	MchID           string
	MchAPIKey       string //商户平台设置的api key
}
