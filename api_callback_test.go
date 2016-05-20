package godingtalk

import (
    "testing"
)

func TestRegisterCallback(t *testing.T) {
    err:=c.RegisterCallback([]string{"user_add_org"},"hello","http://www.dingtalk.com")
    if err!=nil {
        t.Error(err)
    }
}