package godingtalk

import (
	"fmt"
	"net/url"
	"time"
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
	Expires     int    `json:"expires_in"`
	Created     int64
}

func (e *AccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *AccessTokenResponse) ExpiresIn() int {
	return e.Expires
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
	cache := NewFileCache(".auth_file")
	err := cache.Get(&data)
	if err == nil {
		c.accessToken = data.AccessToken
		return nil
	}

	params := url.Values{}
	params.Add("corpid", c.corpID)
	params.Add("corpsecret", c.corpSecret)
	err = c.httpRPC("gettoken", params, nil, &data)
	if err == nil {
		c.accessToken = data.AccessToken
		if err == nil {
			data.Expires = data.Expires | 7200
			data.Created = time.Now().Unix()
			cache.Set(&data)
		}
	}
	return err
}

//GetJsAPITicket is to retrieve ticket for JS API
func (c *DingTalkClient) GetJsAPITicket() (string, error) {
	return "", nil
}
