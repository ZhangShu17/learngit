package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//----------------------------------
//万年历调用代码 -极速数据
//接口地址：http://api.jisuapi.com/calendar/query
//支持格式:JSON
//请求方法：GET
const APPKEY = "056140b55464ff51"                       //appkey
const jisuURL = "http://api.jisuapi.com/calendar/query" //接口地址
const Date = "2017-9-21"                                //查询时间

// test
func main() {
	//获取详细的当天信息，函数
	Request()
}

func Request() {
	//初始化参数
	param := url.Values{}

	//配置请求参数，方法内部已处理urlencode问题
	param.Set("appkey", APPKEY)
	param.Set("date", Date)

	//发送请求
	data, err := Get(jisuURL, param)
	if err != nil {
		fmt.Errorf("请求失败，错误信息：\r\n%v", err)
	} else {
		//定义映射格式的接口数据
		var netReturn map[string]interface{}
		json.Unmarshal(data, &netReturn) //将json格式的数据解析到interface接口类型，并传递给netReturn的地址
		if netReturn["status"] == "0" {
			fmt.Printf("接口返回的msg字段是：\r\n%v", netReturn["msg"])
			fmt.Printf("接口返回的result字段是：\r\v%v", netReturn["result"])
		}
	}
}

//网络请求
func Get(apiURL string, param url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url出错：", err)
		return nil, err
	}
	Url.RawQuery = param.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//jdjksdkdsl
//sjjdjdjdnjhshdjjjjdjs还是
