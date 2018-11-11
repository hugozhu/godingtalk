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

const typeJSON = "application/json"

//UploadFile is for uploading a single file to DingTalk
type UploadFile struct {
	FieldName string
	FileName  string
	Reader    io.Reader
}

//DownloadFile is for downloading a single file from DingTalk
type DownloadFile struct {
	MediaID  string
	FileName string
	Reader   io.Reader
}

func (c *DingTalkClient) httpRPC(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
	if c.AccessToken != "" {
		if params == nil {
			params = url.Values{}
		}
		if params.Get("access_token") == "" {
			params.Set("access_token", c.AccessToken)
		}
	}
	return c.httpRequest(path, params, requestData, responseData)
}

func (c *DingTalkClient) httpRequest(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
	client := c.HTTPClient
	var request *http.Request
	url2 := ROOT + path + "?" + params.Encode()
	// log.Println(url2)
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
			// log.Printf("url: %s request: %s", url2, string(d))
			request, _ = http.NewRequest("POST", url2, bytes.NewReader(d))
			request.Header.Set("Content-Type", typeJSON)
		}
	} else {
		// log.Printf("url: %s", url2)
		request, _ = http.NewRequest("GET", url2, nil)
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Server error: " + resp.Status)
	}

	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	//log.Printf("url: %s response content type: %s", url2, contentType)
	pos := len(typeJSON)
	if len(contentType) >= pos && contentType[0:pos] == typeJSON {
		content, err := ioutil.ReadAll(resp.Body)
		//log.Println(string(content))
		if err == nil {
			json.Unmarshal(content, responseData)
			return responseData.checkError()
		}
	} else {
		io.Copy(responseData.getWriter(), resp.Body)
		return responseData.checkError()
	}
	return err
}
