package godingtalk

import (
	"os"
	"testing"
)

var c *DingTalkClient

func init() {
	c = NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	c.RefreshAccessToken()
}

func TestDepartmentApi(t *testing.T) {
	departments, err := c.DepartmentList()
	// t.Logf("%+v", departments)
	d, err := c.DepartmentDetail(departments.Departments[0].Id)
	if err != nil {
		t.Error(err)
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

func TestMessageApi(t *testing.T) {
	// err := c.SendAppMessage("22194403", "0420506555", "测试消息，请忽略") //@all
	// if err != nil {
	// 	t.Error(err)
	// }
	// err = c.SendTextMessage("0420506555", "chat6a93bc1ee3b7d660d372b1b877a9de62", "测试消息，请忽略")
	// if err != nil {
	// 	t.Error(err)
	// }
}
