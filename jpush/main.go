package main

import (
	"fmt"
	"jpush/api/push"
)

const (
	appKey = "fc4e7512bfd31c3dba9bad3e"
	secret = "570d3442dc88d1382fe93b89"
)

const jsons = `{
   "platform" : "all",
   "audience" : "all",
   "notification" : {
      "alert" : "Hi, go!" 
   }
}`

func main() {

	//PlatForm
	var pf push.Platform
	//pf.Add(push.ANDROID)
	pf.All()

	//Audience
	var ad push.Audience
	//s := []string{"1", "2", "3"}
	//ad.SetID(s)
	ad.All()

	//Notice
	var notice push.AndroidNotice
	notice.SetAlert("alert_test")
	notice.SetTitle("title_test")

	var msg push.Message
	msg.Title = "Hello"
	msg.Content = "祝大家工作顺利"

	//NoticeBuilder
	//nb := push.NewNoticeBuilder()
	//nb.SetPlatform(&pf)
	//nb.SetAudience(&ad)
	//nb.SetSimpleNotice("简单通知")
	//nb.SetAndroidNotice(&notice)

	mb := push.NewMessageBuilder()
	mb.SetPlatform(&pf)
	mb.SetAudience(&ad)
	mb.SetMessage(&msg)

	//push
	c := push.NewPushClient(secret, appKey)
	str, err := c.Send(mb)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", str)
	}

}
