package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	dingtalk "github.com/hugozhu/godingtalk"
)

var msgType string
var agentID string
var senderID string
var toUser string
var chatID string
var content string
var link string

func init() {
	flag.StringVar(&msgType, "type", "app", "message type")
	flag.StringVar(&agentID, "agent", "22194403", "agent Id")
	flag.StringVar(&senderID, "sender", "011217462940", "sender id")
	flag.StringVar(&toUser, "touser", "0420506555", "touser id")
	flag.StringVar(&chatID, "chat", "chat6a93bc1ee3b7d660d372b1b877a9de62", "chat id")
	flag.StringVar(&link, "link", "http://hugozhu.myalert.info/dingtalk", "link url")
	flag.Parse()
}

func usage() {
	flag.Usage()
	os.Exit(-1)
}

func main() {
	c := dingtalk.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	c.RefreshAccessToken()
	var err error
	if len(os.Args) < 2 {
		usage()
	}
	content = os.Args[len(os.Args)-1]
	switch msgType {
	case "app":
		err = c.SendAppMessage(agentID, toUser, content)
	case "text":
		err = c.SendTextMessage(senderID, chatID, content)
	case "oa":
		msg := dingtalk.OAMessage{}
		json.Unmarshal([]byte(content), &msg)
		err = c.SendOAMessage(senderID, chatID, msg)
	}
	if err != nil {
		fmt.Println(err)
	}
}
