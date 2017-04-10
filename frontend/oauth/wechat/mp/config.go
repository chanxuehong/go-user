package mp

import (
	"github.com/chanxuehong/wechat/mp/user/oauth2"

	"github.com/chanxuehong/go-user/config"
)

var oauth2Config = oauth2.NewOAuth2Config(
	config.ConfigData.Weixin.MP.AppId,
	config.ConfigData.Weixin.MP.AppSecret,
	"unused",
	"snsapi_userinfo",
)
