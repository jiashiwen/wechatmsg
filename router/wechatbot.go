package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitWeChatBotRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	//BaseRouter := Router.Group("base").Use(middleware.NeedInit())
	WeChatBotRouter := Router.Group("wechat")
	{
		WeChatBotRouter.POST("sendtext", v1.SendTextToFriend)
		WeChatBotRouter.POST("sendtexttogroup", v1.SendTextToGroup)
		WeChatBotRouter.GET("getgroups", v1.GetWeChatGroups)

	}
	return WeChatBotRouter
}
