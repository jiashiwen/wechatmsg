package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"github.com/eatMoreApple/openwechat"
	"time"
)

func SendTextToFriend(msg request.WeChatSendMsg) error {
	bot := global.GetWeCahtBotServer().Bot

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return err
	}
	friends, _ := self.Friends()
	fr := friends.SearchByNickName(1, msg.NickName)

	fr.SendText(time.Now().String()+" "+msg.Msg, 1*time.Second)
	return nil

}

func SendTextToGroup(msg request.WeChatSendMsg) error {
	// 获取登陆的用户
	self, err := global.GetWeCahtBotServer().Bot.GetCurrentUser()
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return err
	}

	global.GVA_LOG.Sugar().Info("msg:", msg.Msg, "|", "nickname", msg.NickName)
	groups, _ := self.Groups()
	group := groups.SearchByNickName(1, msg.NickName)
	group.SendText(time.Now().String()+"|"+msg.Msg, 1*time.Second)
	return nil
}

func GetFriends() (openwechat.Friends, error) {
	self, err := global.GetWeCahtBotServer().Bot.GetCurrentUser()
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return nil, err
	}
	return self.Friends(true)
}

func GetGroups() (openwechat.Groups, error) {
	self, err := global.GetWeCahtBotServer().Bot.GetCurrentUser()
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return nil, err
	}

	return self.Groups(true)
}

//func StartWeChatBot() {
//
//}
