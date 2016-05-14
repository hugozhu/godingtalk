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

//SendImageMessage is 发送图片消息
func (c *DingTalkClient) SendImageMessage(sender string, cid string, mediaID string) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "image",
		"image": map[string]string{
			"media_id": mediaID,
		},
	}
	err := c.httpRPC("chat/send", nil, request, &data)
	return err
}

//SendVoiceMessage is 发送语音消息
func (c *DingTalkClient) SendVoiceMessage(sender string, cid string, mediaID string, duration string) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "voice",
		"voice": map[string]string{
			"media_id": mediaID,
			"duration": duration,
		},
	}
	err := c.httpRPC("chat/send", nil, request, &data)
	return err
}

//SendFileMessage is 发送文件消息
func (c *DingTalkClient) SendFileMessage(sender string, cid string, mediaID string) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "file",
		"file": map[string]string{
			"media_id": mediaID,
		},
	}
	err := c.httpRPC("chat/send", nil, request, &data)
	return err
}

//SendLinkMessage is 发送链接消息
func (c *DingTalkClient) SendLinkMessage(sender string, cid string, mediaID string, url string, title string, text string) error {
	var data OAPIResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "link",
		"link": map[string]string{
			"messageUrl": url,
			"picUrl":     mediaID,
			"title":      title,
			"text":       text,
		},
	}
	err := c.httpRPC("chat/send", nil, request, &data)
	return err
}

//OAMessage is the Message for OA
type OAMessage struct {
	URL  string `json:"message_url"`
	Head struct {
		BgColor string `json:"bgcolor,omitempty"`
		Text    string `json:"text,omitempty"`
	} `json:"head,omitempty"`
	Body struct {
		Title     string          `json:"title,omitempty"`
		Form      []OAMessageForm `json:"form,omitempty"`
		Rich      OAMessageRich   `json:"rich,omitempty"`
		Content   string          `json:"content,omitempty"`
		Image     string          `json:"image,omitempty"`
		FileCount int             `json:"file_count,omitempty"`
		Author    string          `json:"author,omitempty"`
	} `json:"body,omitempty"`
}

type OAMessageForm struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type OAMessageRich struct {
	Num  string `json:"num,omitempty"`
	Unit string `json:"body,omitempty"`
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
