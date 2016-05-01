package godingtalk

//SendAppMessage is 发送企业会话消息
func (c *DingTalkClient) SendAppMessage(agentID string, msg string) error {
    var data OAPIResponse
    request := map[string]interface{} {
        "touser":"@all",
        "agentid":agentID,
        "msgtype":"text",
        "text": map[string]interface{} {
            "content": msg,
        },
    }
    err :=c.httpRPC("message/send", nil, request, &data)
    return err
}


//SendTextMessage is 发送普通文本消息
func (c *DingTalkClient) SendTextMessage(sender string, cid string, msg string) error {
    var data OAPIResponse
    request := map[string]interface{} {
        "chatid":cid,
        "sender":sender,
        "msgtype":"text",
        "text": map[string]interface{} {
            "content": msg,
        },
    }
    err :=c.httpRPC("chat/send", nil, request, &data)
    return err
}