package godingtalk

import (
	"testing"
)

func TestRegisterCallback(t *testing.T) {
	err := c.UpdateCallback([]string{"user_modify_org"}, "hello", "1234567890123456789012345678901234567890aes", "http://go.myalert.info:8888/dingtalk/callback/")
	if err != nil {
		t.Log(err)
	}
	err = c.DeleteCallback()
	if err != nil {
		t.Log(err)
	}
	err = c.RegisterCallback([]string{"user_add_org"}, "hello", "1234567890123456789012345678901234567890aes", "http://go.myalert.info:8888/dingtalk/callback/")
	if err != nil {
		t.Log(err)
	}
}

func TestListCallback(t *testing.T) {
	data, err := c.ListCallback()
	if err != nil {
		t.Log(err)
	}
	t.Log(data)
}
