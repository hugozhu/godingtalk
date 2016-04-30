package godingtalk

import (
	"testing"
    "os"
)

var c *DingTalkClient

func init() {
    c = NewDingTalkClient(os.Getenv("corpid"),os.Getenv("corpsecret"))
    c.refreshAccessToken()    
}

func TestGET(t *testing.T) {    
    departments, err:=c.DepartmentList()
    d,err := c.DepartmentDetail(departments.Departments[0].Id)
    if err!=nil {
        t.Error(err)
    }    
    if (d.Id != departments.Departments[0].Id) {
        t.Error("DepartmentDetail error")
    }
}

func TestSend(t *testing.T) {
    err := c.SendMessage("22194403", "Hello World")
    if err!=nil {
        t.Error(err)
    }       
}