package global

import "gin-vue-admin/wechatbot"

var wbs *wechatbot.WechatBotServer

// 初始化WechatBotServer
func InitWeChatBotServer() {
	wbs = wechatbot.NewWechatBotServer()
}
