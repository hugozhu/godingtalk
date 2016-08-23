package godingtalk

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	str, err := c.Encrypt("Hello")
    if err!=nil {
        t.Error(err)
    } else {
        t.Log(str)
    }
}
