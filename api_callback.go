package godingtalk

//RegisterCallback is 注册事件回调接口
func (c *DingTalkClient) RegisterCallback(callbacks []string, token string, callbackURL string) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"call_back_tag":  callbacks,
		"token": token,
		"aes_key": "1",
		"url": callbackURL,
	}
	err := c.httpRPC("call_back/register_call_back", nil, request, &data)
	return err
}