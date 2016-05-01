package godingtalk

import (
	"fmt"
	"net/url"
)

const (
	//ROOT is the root url
	ROOT = "https://oapi.dingtalk.com/"
)

//DingTalkClient is the Client to access DingTalk Open API
type DingTalkClient struct {
	corpID      string
	corpSecret  string
	accessToken string
}

//Marshallable is
type Marshallable interface {
	marshal() []byte
}

//Unmarshallable is
type Unmarshallable interface {
	checkError() error
}

//OAPIResponse is
type OAPIResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (data *OAPIResponse) checkError() (err error) {
	if data.ErrCode != 0 {
		err = fmt.Errorf("%d: %s", data.ErrCode, data.ErrMsg)
	}
	return err
}

//AccessTokenResponse is
type AccessTokenResponse struct {
	OAPIResponse
	AccessToken string `json:"access_token"`
}

//NewDingTalkClient creates a DingTalkClient instance
func NewDingTalkClient(corpID string, corpSecret string) *DingTalkClient {
	c := new(DingTalkClient)
	c.corpID = corpID
	c.corpSecret = corpSecret
	return c
}

//RefreshAccessToken is
func (c *DingTalkClient) RefreshAccessToken() error {
	var data AccessTokenResponse
	params := url.Values{}
	params.Add("corpid", c.corpID)
	params.Add("corpsecret", c.corpSecret)
	err := c.httpRPC("gettoken", params, nil, &data)
	if err == nil {
		c.accessToken = data.AccessToken
	}
	return err
}
