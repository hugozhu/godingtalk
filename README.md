# dingtalk-openapi-sdk
DingTalk Open API SDK in golang

http://open.dingtalk.com


## Usage

Fetch the SDK
```
export GOPATH=`pwd`
go get github.com/hugozhu/godingtalk
export corpid=...
export corpsecret=...
```

Example to send a message

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
	err := c.SendMessage("22194403", "此消息通过SDK github.com/hugozhu/godingtalk 发出")
	log.Println(err)
}
```
