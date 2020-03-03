package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	dingtalk "github.com/hugozhu/godingtalk"
)

var msgType string
var agentID string
var senderID string
var toUser string
var chatID string
var content string
var link string
var file string
var title string
var text string
var robot bool
var token string

func init() {
	flag.BoolVar(&robot, "robot", false, "use robot api?")
	flag.StringVar(&token, "token", "", "robot access token or token to override env setting")
	flag.StringVar(&msgType, "type", "app", "message type (app, text, image, voice, link, oa, markdown )")
	flag.StringVar(&agentID, "agent", "22194403", "agent Id")
	flag.StringVar(&senderID, "sender", "011217462940", "sender id")
	flag.StringVar(&toUser, "touser", "0420506555", "touser id")
	flag.StringVar(&chatID, "chat", "chat6a93bc1ee3b7d660d372b1b877a9de62", "chat id")
	flag.StringVar(&link, "link", "http://hugozhu.myalert.info/dingtalk", "link url")
	flag.StringVar(&file, "file", "", "file path for media message")
	flag.StringVar(&title, "title", "This is link title", "title for link message")
	flag.StringVar(&text, "text", "This is link text", "text for link message")
	flag.Parse()
}

func usage() {
	flag.Usage()
	os.Exit(-1)
}

func fatalError(err error) {
	fmt.Println(err)
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
		if robot {
			_, err = c.SendRobotTextMessage(token, content)
		} else {
			_, err = c.SendTextMessage(senderID, chatID, content)
		}
	case "markdown":
		c.SendRobotMarkdownMessage(token, title, content)
	case "image":
		if file == "" {
			panic("Image path is empty")
		}
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			fatalError(err)
		}
		media, err := c.UploadMedia("image", filepath.Base(file), f)
		if err != nil {
			fatalError(err)
		}
		_, err = c.SendImageMessage(senderID, chatID, media.MediaID)
		if err != nil {
			fatalError(err)
		}
	case "voice":
		if file == "" {
			panic("Voice file path is empty")
		}
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			fatalError(err)
		}
		media, err := c.UploadMedia("voice", filepath.Base(file), f)
		if err != nil {
			fatalError(err)
		}
		_, err = c.SendVoiceMessage(senderID, chatID, media.MediaID, "10")
		if err != nil {
			fatalError(err)
		}
	case "file":
		if file == "" {
			panic("File path is empty")
		}
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			fatalError(err)
		}
		media, err := c.UploadMedia("file", filepath.Base(file), f)
		if err != nil {
			fatalError(err)
		}
		_, err = c.SendFileMessage(senderID, chatID, media.MediaID)
		if err != nil {
			fatalError(err)
		}
	case "link":
		if file == "" {
			panic("File path is empty")
		}
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			fatalError(err)
		}
		media, err := c.UploadMedia("image", filepath.Base(file), f)
		if err != nil {
			fatalError(err)
		}
		_, err = c.SendLinkMessage(senderID, chatID, media.MediaID, link, title, text)
		if err != nil {
			fatalError(err)
		}
	case "oa":
		msg := dingtalk.OAMessage{}
		json.Unmarshal([]byte(content), &msg)
		_, err = c.SendOAMessage(senderID, chatID, msg)
	}
	if err != nil {
		fmt.Println(err)
	}
}
