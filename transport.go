package godingtalk

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/json"
)

func (c *DingTalkClient) httpRPC(path string, params url.Values, data Unmarshallable) error {
    client := new (http.Client)    
    if (c.accessToken != "") { 
        if (params==nil) {
            params = url.Values{}
        }       
        params.Set("access_token", c.accessToken)        
    }    
    url2 := ROOT + path + "?"+ params.Encode()    
    resp, err := client.Get(url2)
    if err != nil {        
        return err
    } 
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)    
    if err == nil {
        json.Unmarshal(content, data)
        err = data.checkError()
    }
    return err        
}