package main

import (
	"fmt"
	"github.com/ipandtcp/godingtalk"
	"os"
)

func main() {

	c := godingtalk.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	err := c.RefreshAccessToken()
	if err != nil {
		panic(err)
	}

	c.AgentID = "161271936"
	msg := godingtalk.OAMessage{}
	msg.URL = "http://www.google.com/"
	msg.Head.Text = "头部标题"
	msg.Head.BgColor = "FFBBBBBB"
	msg.Body.Title = "正文标题"
	msg.Body.Content = "test content"
	taskID, err := c.TopAPIMsgSend("oa", []string{"085354234826136236"}, nil, false, msg)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("taskID: %d\n", taskID)
	}

	sendProgress, err := c.TopAPIMsgGetSendProgress(taskID)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%v\n", sendProgress.OK.Progress)
	}

	sendResult, err := c.TopAPIMsgGetSendResult(taskID)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%v\n", sendResult.OK.SendResult)
	}
}
