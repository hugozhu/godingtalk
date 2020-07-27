package godingtalk

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	str, err := c.Encrypt("Hello")
    if err!=nil {
        t.Log(err)
    } else {
        t.Log(str)
    }
}
