package godingtalk

import (
	"os"
	"testing"
	"time"
)

var c *DingTalkClient

func init() {
	c = NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	err := c.RefreshAccessToken()
	if err != nil {
		panic(err)
	}
}

func TestInitWithAppKey(t *testing.T) {
	c1 := NewDingTalkClient(os.Getenv("appkey"), os.Getenv("appsecret"))
	err := c1.RefreshAccessToken(true)
	if err != nil {
		panic(err)
	}
	_, err = c1.SendRobotTextMessage(os.Getenv("token"), "Message sent successfully with appkey and appsecret")
	if err != nil {
		t.Error(err)
	}
}

func TestCalendarListApi(t *testing.T) {
	from := time.Now().AddDate(0, 0, -1)
	to := time.Now().AddDate(0, 0, 1)
	_, err := c.ListEvents("0420506555", from, to)
	if err != nil {
		panic(err)
	}
	//for _, event := range events {
	//	t.Logf("%v %v %v %v", event.Start, event.End, event.Summary, event.Description)
	//}
}

func TestDepartmentApi(t *testing.T) {
	departments, err := c.DepartmentList()
	// t.Logf("%+v %+v", departments, err)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	d, err := c.DepartmentDetail(departments.Departments[0].Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if d.Id != departments.Departments[0].Id {
		t.Error("DepartmentDetail error")
	}

	for _, department := range departments.Departments {
		t.Logf("dept: %v", department)
		list, err := c.UserList(department.Id, 0, 100)
		if err != nil {
			t.Error(err)
		}
		for _, user := range list.Userlist {
			t.Logf("\t\tuser: %v", user)
		}
	}
}

func TestJsAPITicket(t *testing.T) {
	ticket, err := c.GetJsAPITicket()
	if err != nil || ticket == "" {
		t.Error("JsAPITicket error", err)
	}
}

func TestCreateChat(t *testing.T) {
	// chatid, err := c.CreateChat("Test chat", "0420506555", []string{"0420506555"})
	// if err!=nil {
	//     t.Error(err)
	// }
	// t.Log("-----",chatid)
}

func TestSendAppMessageApi(t *testing.T) {
	err := c.SendAppMessage("22194403", "0420506555", "测试消息，请忽略") //@all
	if err != nil {
		t.Error(err)
	}
}

func TestTextMessage(t *testing.T) {
	data, err := c.SendTextMessage("011217462940", "chat6a93bc1ee3b7d660d372b1b877a9de62", "测试消息，来自双十一，请忽略")
	if err != nil {
		t.Error(err)
	} else {
		if data.MessageID == "" {
			t.Error("Message id is empty")
		}
	}
	data2, _ := c.GetMessageReadList(data.MessageID, 0, 10)
	if len(data2.ReadUserIdList) == 0 {
		t.Error("Message Read List should not be empty")
	}
}

func TestSendOAMessage(t *testing.T) {
	msg := OAMessage{}
	msg.URL = "http://www.google.com/"
	msg.Head.Text = "头部标题"
	msg.Head.BgColor = "FFBBBBBB"
	msg.Body.Title = "正文标题"
	msg.Body.Content = "test content"
	_, err := c.SendOAMessage("011217462940", "chat6a93bc1ee3b7d660d372b1b877a9de62", msg)
	if err != nil {
		t.Error(err)
	}
}

func _TestDownloadAndUploadImage(t *testing.T) {
	f, err := os.Create("lADOHrf_oVxc.jpg")
	if err == nil {
		err = c.DownloadMedia("@lADOHrf_oVxc", f)
	}
	if err != nil {
		t.Error(err)
	}
	f.Close()

	f, _ = os.Open("lADOHrf_oVxc.jpg")
	defer f.Close()
	media, err := c.UploadMedia("image", "myfile.jpg", f)
	if media.MediaID == "" {
		t.Error("Upload File Failed")
	}
	t.Log("uploaded file mediaid:", media.MediaID)
	if err != nil {
		t.Error(err)
	}
	_, err = c.SendImageMessage("011217462940", "chat6a93bc1ee3b7d660d372b1b877a9de62", "@lADOHrf_oVxc")
	if err != nil {
		t.Error(err)
	}
}

func TestVoiceMessage(t *testing.T) {
	// f, _ := os.Open("/Users/hugozhu/Downloads/BlackBerry_test2_AMR-NB_Mono_12.2kbps_8000Hz.amr")
	// defer f.Close()
	// media, err := c.UploadMedia("voice", "sample.amr", f)
	// if media.MediaID == "" {
	// 	t.Error("Upload File Failed")
	// }
	// t.Log("uploaded file mediaid:", media.MediaID)
	// if err != nil {
	// 	t.Error(err)
	// }
	_, err := c.SendVoiceMessage("011217462940", "chat6a93bc1ee3b7d660d372b1b877a9de62", "@lATOHr53E84DALnDzml4wS0", "10")
	if err != nil {
		t.Error(err)
	}
}

func TestRobotMessage(t *testing.T) {
	_, err := c.SendRobotTextMessage(os.Getenv("token"), "这是一条测试消息")
	if err != nil {
		t.Error(err)
	}

	_, err = c.SendRobotMarkdownMessage(os.Getenv("token"), "测试标题", "# 杭州天气\n这是一条测试消息\n"+
		"> 9度，西北风1级，空气良89，**相对温度**73%\n\n"+
		"> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n"+
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	if err != nil {
		t.Error(err)
	}
}

func TestRobotAtMessage(t *testing.T) {
	_, err := c.SendRobotTextAtMessage(os.Getenv("token"), "这是一条测试消息", &RobotAtList{
		IsAtAll: true,
	})
	if err != nil {
		t.Error(err)
	}
}
