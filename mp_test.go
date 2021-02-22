package gowechat

import (
	"github.com/kprince/gowechat/wxcontext"
	"io/ioutil"
	"log"
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

func TestWechat_Mini(t *testing.T) {
	config := wxcontext.Config{
		MiniAppID:     "",
		MiniAppSecret: "",
		Token:     "",
	}
	wc := NewWechat(config)
	mi, _ := wc.Mini()
	bts,_ :=mi.GetMiniCode("test",512,true)
	log.Println(string(bts))
	if len(bts)>0 {
		ioutil.WriteFile("1.jpg",bts,0755)
	}
}
