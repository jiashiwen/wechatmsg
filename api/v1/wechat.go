package v1

import (
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

func SendTextToFriend(c *gin.Context) {
	var msg request.WeChatSendMsg
	_ = c.ShouldBindJSON(&msg)
	if err := service.SendTextToFriend(msg); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

func SendTextToGroup(c *gin.Context) {
	var msg request.WeChatSendMsg
	_ = c.ShouldBindJSON(&msg)
	if err := service.SendTextToGroup(msg); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

func GetWeChatGroups(c *gin.Context) {
	nicknames := []string{}
	groups, err := service.GetGroups()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	for _, v := range groups {
		nicknames = append(nicknames, v.NickName)

	}

	response.OkWithData(nicknames, c)
}
