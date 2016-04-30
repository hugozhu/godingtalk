package godingtalk

//SendMessage is 发送企业会话消息
func (c *DingTalkClient) SendMessage(agentID string, msg string) error {
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