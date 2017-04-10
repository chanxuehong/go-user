package app

import (
	"github.com/chanxuehong/wechat/open/oauth2"

	"github.com/chanxuehong/go-user/config"
)

var oauth2Config = oauth2.NewConfig(
	config.ConfigData.Weixin.Open.App.AppId,
	config.ConfigData.Weixin.Open.App.AppSecret,
	"unused",
	"snsapi_userinfo",
)
