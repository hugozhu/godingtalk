package godingtalk

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/json"
    "bytes"
)

func (c *DingTalkClient) httpRPC(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
    client := new (http.Client)    
    var request *http.Request
    if (c.accessToken != "") { 
        if (params==nil) {
            params = url.Values{}
        }       
        params.Set("access_token", c.accessToken)        
    }    
    url2 := ROOT + path + "?"+ params.Encode()
    if requestData!=nil {
        d,_ := json.Marshal(requestData)        
        request, _ = http.NewRequest("POST", url2, bytes.NewReader(d))
        request.Header.Set("Content-Type", "application/json")
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
        json.Unmarshal(content, responseData)
        err = responseData.checkError()
    }
    return err        
}