package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// POSTHTTP post 请求返回 json
func POSTHTTP(url string, dataBytes []byte, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(dataBytes))
	if err != nil {
		return []byte(""), err
	}

	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

func POSTJson(path string, postData map[string]interface{}, header map[string]string) ([]byte, error) {

	jsonStr, err := json.Marshal(postData)
	if err != nil {
		return []byte(""), err
	}

	payload := strings.NewReader(string(jsonStr))
	req, _ := http.NewRequest("POST", path, payload)
	req.Header.Add("content-type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}
	/*
	   req.Header.Add("isupdate", isupdate)
	   req.Header.Add("Api-Key", apiKey)
	   req.Header.Add("accept", "application/json")
	*/

	resp, err := HttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

func GET(path string, header map[string]string) ([]byte, error) {

	req, _ := http.NewRequest("GET", path, nil)

	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

//适用于 application/x-www-form-urlencoded
func PostUrlEncoded(path string, postData map[string]string, header map[string]string) ([]byte, error) {
	//post要提交的数据
	DataUrlVal := url.Values{}
	for key, val := range postData {
		DataUrlVal.Add(key, val)
	}
	req, err := http.NewRequest("POST", path, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for key, value := range header {
		req.Header.Add(key, value)
	}

	//提交请求
	resp, err := HttpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	//读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取短链接
func GetShort(timeout time.Duration) *http.Client {
	httpProxy := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: (timeout) * time.Second,
	}
	return httpProxy
}

func PostFormByClient(url string, postData map[string]interface{}, header map[string]string, httpClient *http.Client) ([]byte, error) {
	jsonStr, err := json.Marshal(postData)
	if err != nil {
		return []byte(""), err
	}
	payload := strings.NewReader(string(jsonStr))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Add("content-type", "application/json")
	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}
