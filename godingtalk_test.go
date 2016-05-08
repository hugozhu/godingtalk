package godingtalk

import (
	"os"
	"testing"
)

var c *DingTalkClient

func init() {
	c = NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	err := c.RefreshAccessToken()
	if err != nil {
		panic(err)
	}
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
		list, err := c.UserList(department.Id)
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
	err := c.SendTextMessage("011217462940", "chat6a93bc1ee3b7d660d372b1b877a9de62", "测试消息，请忽略")
	if err != nil {
		t.Error(err)
	}
}

func TestSendOAMessage(t *testing.T) {
	msg := OAMessage{}
	msg.URL = "http://www.google.com/"
	msg.Head.Text = "头部标题"
	msg.Head.BgColor = "FFBBBBBB"
	msg.Body.Title = "正文标题"
	msg.Body.Content = "test content"
	err := c.SendOAMessage("011217462940", "chat6a93bc1ee3b7d660d372b1b877a9de62", msg)
	if err != nil {
		t.Error(err)
	}
}
