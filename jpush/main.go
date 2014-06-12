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
	var pf push.PlatForm
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
	msg.Content = " 祝大家工作顺利55"

	//NoticeBuilder
	//var nb = push.NewNoticeBuilder()
	//nb.SetPlatForm(&pf)
	//nb.SetAudience(&ad)
	//nb.SetAndroidNotice(&notice)
	//nb.SetSimpleNotice("哈哈")

	mb := push.NewMessageBuilder()
	mb.SetPlatForm(&pf)
	mb.SetAudience(&ad)
	mb.SetMessage(&msg)

	//to json
	//body, _ := json.Marshal(mb)
	//fmt.Printf("ok : %s", string(body))
	////push
	c := push.NewPushClient(secret, appKey)
	str, err := c.Send(mb)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", str)
	}

}
