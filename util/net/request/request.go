package util_net_request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	u "net/url"
)

func Get[T any](url string, params map[string]interface{}, obj *T) {
	ps := u.Values{}
	Url, _ := u.Parse(url)

	if params != nil {
		for k, v := range params {
			vStr := fmt.Sprintf("%v", v)
			ps.Set(k, vStr)
		}
	}

	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = ps.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath) //等同于https://www.xxx.com?age=23&name=zhaofan
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	json.Unmarshal(body, obj)
}

func PostForm[T any](url string, params map[string]interface{}, obj *T) {
	ps := u.Values{}

	if params != nil {
		for k, v := range params {
			vStr := fmt.Sprintf("%v", v)
			ps.Add(k, vStr)
		}
	}

	resp, _ := http.PostForm(url, ps)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	json.Unmarshal(body, obj)
}

func PostJson[T any](url string, params map[string]interface{}, obj *T) {
	client := &http.Client{}
	bytesData, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	json.Unmarshal(body, obj)
}
