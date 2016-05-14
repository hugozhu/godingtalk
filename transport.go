package godingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

//UploadFile is for uploading a single file to DingTalk
type UploadFile struct {
	FieldName string
	FileName  string
	Reader    io.Reader
}

func (c *DingTalkClient) httpRPC(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
	client := c.HTTPClient
	var request *http.Request
	if c.AccessToken != "" {
		if params == nil {
			params = url.Values{}
		}
		params.Set("access_token", c.AccessToken)
	}
	url2 := ROOT + path + "?" + params.Encode()
	if requestData != nil {
		switch requestData.(type) {
		case UploadFile:
			var b bytes.Buffer
			request, _ = http.NewRequest("POST", url2, &b)
			w := multipart.NewWriter(&b)

			uploadFile := requestData.(UploadFile)
			if uploadFile.Reader == nil {
				return errors.New("upload file is empty")
			}
			fw, err := w.CreateFormFile(uploadFile.FieldName, uploadFile.FileName)
			if err != nil {
				return err
			}
			if _, err = io.Copy(fw, uploadFile.Reader); err != nil {
				return err
			}
			if err = w.Close(); err != nil {
				return err
			}
			request.Header.Set("Content-Type", w.FormDataContentType())
		default:
			d, _ := json.Marshal(requestData)
			request, _ = http.NewRequest("POST", url2, bytes.NewReader(d))
			request.Header.Set("Content-Type", "application/json")
		}
	} else {
		request, _ = http.NewRequest("GET", url2, nil)
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		// log.Println("response:", string(content))
		json.Unmarshal(content, responseData)
		err = responseData.checkError()
	}
	return err
}
