package godingtalk

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	file, err := c.CreateFile(1024)
	t.Log(file, err)
}
