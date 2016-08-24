package godingtalk

import "testing"
import "time"

type ExpiresData struct {
	Data    string
	Expires int   `json:"expires_in"`
	Created int64 `json:"created"`
}

func (e *ExpiresData) CreatedAt() int64 {
	return e.Created
}

func (e *ExpiresData) ExpiresIn() int {
	return e.Expires
}

func TestFileCache(t *testing.T) {
	cache := NewFileCache(".test_cache")
	data := ExpiresData{
		Data:    "Hello World!",
		Expires: 7200,
		Created: time.Now().Unix(),
	}
	cache.Set(&data)

	var data2 ExpiresData
	cache.Get(&data2)
	t.Logf("%+v %+v", data, data2)

	if data2.Created != data.Created {
		t.Errorf("FileCache error")
	}

	data = ExpiresData{
		Data:    "Hello World 2!",
		Expires: 0,
		Created: time.Now().Unix(),
	}
	cache.Set(&data)
	err := cache.Get(&data2)
	if err == nil {
		t.Error("FileCache error: err should not be nil")
	}
	t.Logf("%+v %+v", data, data2)
}

func TestInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache()
	data := ExpiresData{
		Data:    "Hello World!",
		Expires: 7200,
		Created: time.Now().Unix(),
	}
	cache.Set(&data)

	var data2 ExpiresData
	cache.Get(&data2)
	t.Logf("%+v %+v", data, data2)

	if data2.Created != data.Created {
		t.Errorf("InMemoryCache error")
	}

	data = ExpiresData{
		Data:    "Hello World 2!",
		Expires: 0,
		Created: time.Now().Unix(),
	}
	cache.Set(&data)
	err := cache.Get(&data2)
	if err == nil {
		t.Error("InMemoryCache error: err should not be nil")
	}
	t.Logf("%+v %+v", data, data2)
}
