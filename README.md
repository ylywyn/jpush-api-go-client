jpush-api-go-client
===================

概述
----------------------------------- 
这是JPush REST API 的 go 版本封装开发包,仅支持最新的REST API v3功能。 REST API 文档：http://docs.jpush.cn/display/dev/Push-API-v3
  

使用
----------------------------------- 
go get github.com/swordkee/jpush-api-go-client


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
	notice.SetQuickAppNotice(&QuickAppNotice{Alert: "QuickAppNotice",Title: "test",Page: "url"})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})
      
    //jpushclient.Message
    var msg jpushclient.Message
	msg.Title = "Hello"
	msg.MsgContent = "test"

### 4.构建jpushclient.PayLoad

    req := NewPushRequest()
	req.SetPlatform(&pf)
	req.SetAudience(&ad)
	req.SetMessage(&msg)
	req.SetNotice(&notice)
	req.SetOptions(&op)

### 5.构建PushClient，发出推送

	client := jpushclient.NewClient(secret, appKey)
	result, err := client.Push(req)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", r)
	}

### 6.完整demo

    package main

	import (
		"fmt"
		"github.com/swordkee/jpush-api-go-client"
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
            pf.Add(jpushclient.QUICKAPP)
		pf.Add(jpushclient.WINPHONE)
		//pf.All()

		//Audience
		var ad jpushclient.Audience
		s := []string{"1", "2", "3"}
		ad.SetTag(s)
            ad.SetTagAnd(s)
		ad.SetAlias(s)
		ad.SetID(s)
		//ad.All()

		//Notice
		var notice jpushclient.Notice
		notice.SetAlert("alert_test")
		notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
		notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
            notice.SetQuickAppNotice(&QuickAppNotice{Alert: "QuickAppNotice",Title: "test",Page: "url"})
		notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})

		var msg jpushclient.Message
		msg.Title = "Hello"
		msg.MsgContent = "test"

		req := NewPushRequest()
            req.SetPlatform(&pf)
	        req.SetAudience(&ad)
	        req.SetMessage(&msg)
	        req.SetNotice(&notice)
            req.SetOptions(&op)

		bytes, _ := req.ToBytes()
		fmt.Printf("%s\r\n", string(bytes))

		//push
		c := jpushclient.NewClient(secret, appKey)
		result, err := c.Push(req)
		if err != nil {
			fmt.Printf("err:%s", err.Error())
		} else {
			fmt.Printf("ok:%s", result)
		}
	}

