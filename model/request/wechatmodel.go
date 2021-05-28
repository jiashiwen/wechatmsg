package request

type WeChatSendMsg struct {
	Msg      string `json:"msg"`
	NickName string `json:"nickname"`
}
