# DingTalk Open API golang SDK

![image](http://static.dingtalk.com/media/lALOAQ6nfSvM5Q_229_43.png)

Check out DingTalk Open API document at: http://open.dingtalk.com

## Usage

Fetch the SDK
```
export GOPATH=`pwd`
go get github.com/hugozhu/godingtalk
```

### Example code to send a micro app message

```
package main

import (
    "github.com/hugozhu/godingtalk"
    "log"
    "os"
)

func main() {
    c := godingtalk.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
    c.RefreshAccessToken()
    err := c.SendAppMessage(os.Args[1], os.Args[2], os.Args[3])
    if err != nil {
        log.Println(err)
    }
}
```


## Guide

Step-by-step Guide to use this SDK

http://hugozhu.myalert.info/2016/05/02/66-use-dingtalk-golang-sdk-to-send-message-on-pi.html

## Tools

ding_alert: Command line tool to send app/text/oa ... messages

```
export GOPATH=`pwd`
go get github.com/hugozhu/godingtalk/ding_alert

export corpid=<组织的corpid 通过 https://oa.dingtalk.com 获取>
export corpsecret=<组织的corpsecret 通过 https://oa.dingtalk.com 获取>

./bin/ding_alert
Usage of ./bin/ding_alert:
  -agent string
    	agent Id (default "22194403")
  -chat string
    	chat id (default "chat6a93bc1ee3b7d660d372b1b877a9de62")
  -file string
    	file path for media message
  -link string
    	link url (default "http://hugozhu.myalert.info/dingtalk")
  -sender string
    	sender id (default "011217462940")
  -text string
    	text for link message (default "This is link text")
  -title string
    	title for link message (default "This is link title")
  -touser string
    	touser id (default "0420506555")
  -type string
    	message type (app, text, image, voice, link, oa) (default "app")

```