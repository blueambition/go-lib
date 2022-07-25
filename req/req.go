package req

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Post请求
func Post(reqUrl string, proxy string, header map[string]string, data string) ([]byte, int, error) {
	client := &http.Client{}
	if proxy != "" {
		var proxyURI, _ = url.Parse(proxy)
		client.Transport = &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxyURI),
		}
	}
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(data))
	if err != nil {
		return nil, 0, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, resp.StatusCode, err
		}
		return body, resp.StatusCode, nil
	}
	return nil, resp.StatusCode, errors.New("请求有误")
}

//Get请求
func Get(reqUrl, proxy string, header map[string]string) ([]byte, int, error) {
	client := &http.Client{}
	if proxy != "" {
		var proxyURI, _ = url.Parse(proxy)
		client.Transport = &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxyURI),
		}
	}
	//提交请求
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, resp.StatusCode, err
		}
		return body, resp.StatusCode, nil
	}
	return nil, resp.StatusCode, errors.New("请求有误")
}
