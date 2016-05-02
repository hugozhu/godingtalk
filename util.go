package godingtalk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

type Expirable interface {
	CreatedAt() int64
	ExpiresIn() int
}

type FileCache struct {
	Path string
}

func NewFileCache(path string) FileCache {
	return FileCache{
		Path: path,
	}
}

func (c *FileCache) Set(data Expirable) error {
	bytes, err := json.Marshal(data)
	if err == nil {
		ioutil.WriteFile(c.Path, bytes, 0644)
	}
	return err
}

func (c *FileCache) Get(data Expirable) error {
	bytes, err := ioutil.ReadFile(c.Path)
	if err == nil {
		err = json.Unmarshal(bytes, data)
		if err == nil {
			created := data.CreatedAt()
			expires := data.ExpiresIn()
			if err == nil && time.Now().Unix() > created+int64(expires-60) {
				err = errors.New("Data is already expired")
			}
		}
	}
	return err
}
