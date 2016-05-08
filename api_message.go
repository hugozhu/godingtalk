package godingtalk

//SendAppMessage is 发送企业会话消息
func (c *DingTalkClient) SendAppMessage(agentID string, touser string, msg string) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"touser":  touser,
		"agentid": agentID,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	err := c.httpRPC("message/send", nil, request, &data)
	return err
}

//SendTextMessage is 发送普通文本消息
func (c *DingTalkClient) SendTextMessage(sender string, cid string, msg string) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	err := c.httpRPC("chat/send", nil, request, &data)
	return err
}

type OAMessage struct {
	Url  string `json:"message_url"`
	Head struct {
		BgColor string `json:"bgcolor,omitempty"`
		Text    string `json:"text,omitempty"`
	} `json:"head,omitempty"`
	Body struct {
		Title string `json:"title,omitempty"`
		Form  []struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"body,omitempty"`
		Rich struct {
			Num  string `json:"num,omitempty"`
			Unit string `json:"body,omitempty"`
		} `json:"rich,omitempty"`
	} `json:"body,omitempty"`
	Content   string `json:"content,omitempty"`
	Image     string `json:"image,omitempty"`
	FileCount int    `json:"file_count,omitempty"`
	Author    string `json:"author,omitempty"`
}

//SendOAMessage is 发送OA消息
func (c *DingTalkClient) SendOAMessage(sender string, cid string, msg OAMessage) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "oa",
		"oa":      msg,
	}
	err := c.httpRPC("chat/send", nil, request, &data)
	return err
}
