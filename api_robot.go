package godingtalk

import (
	"net/url"
)

//SendRobotTextMessage can send a text message to a group chat
func (c *DingTalkClient) SendRobotTextMessage(accessToken string, msg string) (data MessageResponse, err error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}
