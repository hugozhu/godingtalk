package godingtalk

import (
	"testing"
)

func TestTopAPImsg(t *testing.T) {
	c.AgentID = "161271936"
	msg := OAMessage{}
	msg.URL = "http://www.google.com/"
	msg.Head.Text = "头部标题"
	msg.Head.BgColor = "FFBBBBBB"
	msg.Body.Title = "正文标题"
	msg.Body.Content = "test content"
	taskID, err := c.TopAPIMsgSend("oa", []string{"085354234826136236"}, nil, false, msg)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%d\n", taskID)
	}

	sendProgress, err := c.TopAPIMsgGetSendProgress(taskID)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v\n", sendProgress.OK.Progress)
	}

	sendResult, err := c.TopAPIMsgGetSendResult(taskID)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v\n", sendResult.OK.SendResult)
	}
}
