package req

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var URI, _ = url.Parse("http://127.0.0.1:7890") //本地测试翻墙设置

//Post请求
func Post(reqUrl string, setProxy bool, header map[string]string, postData string) ([]byte, error) {
	client := &http.Client{}
	if setProxy {
		client.Transport = &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(URI),
		}
	}
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(postData))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	//post请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	return nil, errors.New("请求有误")
}

//Get请求
func Get(reqUrl string, setProxy bool, header map[string]string) ([]byte, error) {
	client := &http.Client{}
	if setProxy {
		client.Transport = &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(URI),
		}
	}
	//提交请求
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	return nil, errors.New("请求有误")
}
