jpush-api-go-client
===================

概述
----------------------------------- 
   这是JPush REST API 的 go 版本封装开发包,仅支持最新的REST API v3功能。
   REST API 文档：http://docs.jpush.cn/display/dev/Push-API-v3
  

使用  
----------------------------------- 
   go get github.com/ylywyn/jpush-api-go-client
   
   
推送流程  
----------------------------------- 
### 1.构建要推送的平台： jpushclient.Platform
	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	pf.Add(jpushclient.WINPHONE)
	//pf.All()
      
### 2.构建接收听众： jpushclient.Audience
	//Audience
	var ad jpushclient.Audience
	s := []string{"t1", "t2", "t3"}
	ad.SetTag(s)
	id := []string{"1", "2", "3"}
	ad.SetID(id)
	//ad.All()
      
### 3.构建通知 jpushclient.Notice，或者消息： jpushclient.Message
      
	//Notice
	var notice jpushclient.Notice
	notice.SetAlert("alert_test")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})
      
    //jpushclient.Message
    var msg jpushclient.Message
	msg.Title = "Hello"
	msg.Content = "你是ylywn"
      
### 4.构建jpushclient.PayLoad
    payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)
      
      
### 5.构建PushClient，发出推送
	c := jpushclient.NewPushClient(secret, appKey)
	r, err := c.Send(bytes)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", r)
	}

  
### 6.完整demo
    package main

	import (
		"fmt"
		"github.com/ylywyn/jpush-api-go-client"
	)

	const (
		appKey = "you jpush appkey"
		secret = "you jpush secret"
	)

	func main() {

		//Platform
		var pf jpushclient.Platform
		pf.Add(jpushclient.ANDROID)
		pf.Add(jpushclient.IOS)
		pf.Add(jpushclient.WINPHONE)
		//pf.All()

		//Audience
		var ad jpushclient.Audience
		s := []string{"1", "2", "3"}
		ad.SetTag(s)
		ad.SetAlias(s)
		ad.SetID(s)
		//ad.All()

		//Notice
		var notice jpushclient.Notice
		notice.SetAlert("alert_test")
		notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
		notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
		notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})

		var msg jpushclient.Message
		msg.Title = "Hello"
		msg.Content = "你是ylywn"

		payload := jpushclient.NewPushPayLoad()
		payload.SetPlatform(&pf)
		payload.SetAudience(&ad)
		payload.SetMessage(&msg)
		payload.SetNotice(&notice)

		bytes, _ := payload.ToBytes()
		fmt.Printf("%s\r\n", string(bytes))

		//push
		c := jpushclient.NewPushClient(secret, appKey)
		str, err := c.Send(bytes)
		if err != nil {
			fmt.Printf("err:%s", err.Error())
		} else {
			fmt.Printf("ok:%s", str)
		}
	}

