package global

import (
	"gin-vue-admin/config"
	"gin-vue-admin/wechatbot"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
	once    sync.Once
)

//获取单例wechartbot 单例
func GetWeCahtBotServer() *wechatbot.WechatBotServer {
	once.Do(func() {
		InitWeChatBotServer()
	})
	return wbs
}
