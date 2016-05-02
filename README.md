# DingTalk Open API golang SDK

![image](http://static.dingtalk.com/media/lALOAQ6nfSvM5Q_229_43.png)

Check out DingTalk Open API document at: http://open.dingtalk.com

## Usage

Fetch the SDK
```
export GOPATH=`pwd`
go get github.com/hugozhu/godingtalk
```

### Example code to send an micro app message

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


###  Run the example code
```
export corpid=<组织的corpid 通过 https://oa.dingtalk.com 获取>
export corpsecret=<组织的corpsecret 通过 https://oa.dingtalk.com 获取>

go run src/main.go <agentid 通过 https://oa.dingtalk.com 获取> <userid 或 @all> "消息内容"
```