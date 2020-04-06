package gowechat

import (
	"kprince/gowechat/wxcontext"
	"testing"

	"github.com/astaxie/beego"
)

func TestGetQrcode(t *testing.T) {
	config := wxcontext.Config{
		AppID:     "wx88c493c9a9f67ea6",
		AppSecret: "0c1fe2db856c60d9c52b65383feadae1",
		Token:     "zyt2864",
	}
	wc := NewWechat(config)
	beego.Debug("wechat's cache:", wc.Context.Cache)
	_, _ = wc.MchMgr()
}
