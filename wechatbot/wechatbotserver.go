package wechatbot

import (
	"context"
	"fmt"
	"github.com/eatMoreApple/openwechat"
	"sync"
)

type WechatBotServer struct {
	Bot        *openwechat.Bot
	BotContext context.Context
	BotCancel  context.CancelFunc
	Started    bool
}

func NewWechatBotServer() *WechatBotServer {
	ctx, cancel := context.WithCancel(context.Background())
	ws := WechatBotServer{
		Bot:        openwechat.DefaultBot(),
		BotContext: ctx,
		BotCancel:  cancel,
		Started:    false,
	}
	return &ws
}

//启动微信机器人
func (wb *WechatBotServer) StartBotServer(wg *sync.WaitGroup) error {
	defer wg.Done()
	defer wb.BotCancel()

	// 注册消息处理函数
	wb.Bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	wb.Bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	//if err := wb.Bot.Login(); err != nil {
	//	return err
	//}

	//热登录
	storage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	if err := wb.Bot.HotLogin(storage); err != nil {

		if err := wb.Bot.Login(); err != nil {
			return err
		}

	}

	// 获取登陆的用户
	self, err := wb.Bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 获取所有的群组
	groups, _ := self.Groups()
	for _, v := range groups {
		fmt.Println(v)
	}

	//消息回调
	messageHandle := func(msg *openwechat.Message) {
		if msg.IsText() {
			from, _ := msg.Sender()
			fmt.Println("你收到了一条新的文本消息", msg.Content, "|", from.NickName, "|", from.UserName)

		}
	}
	wb.Bot.MessageHandler = messageHandle

	// 阻塞主goroutine, 知道发生异常或者用户主动退出
	wb.Started = true
	wb.Bot.Block()
	return nil

}

//服务停止
func (wb *WechatBotServer) StopBotServer() {
	wb.Bot.Logout()
	wb.Started = false
	wb.BotCancel()
}

//获取当前用户
func (wb *WechatBotServer) GetSelf() (*openwechat.Self, error) {
	return wb.Bot.GetCurrentUser()
}
