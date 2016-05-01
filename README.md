# dingtalk-openapi-golang-sdk

DingTalk Open API golang SDK

![image](http://static.dingtalk.com/media/lALOAQ6nfSvM5Q_229_43.png)
Checkout DingTalk Open API document at: http://open.dingtalk.com

## Usage

Fetch the SDK
```
export GOPATH=`pwd`
go get github.com/hugozhu/godingtalk
export corpid=<组织的corpid 通过 https://oa.dingtalk.com 获取>
export corpsecret=<组织的corpsecret 通过 https://oa.dingtalk.com 获取>
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
