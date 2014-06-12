jpush-api-go-client
===================

概述
----------------------------------- 
   这是JPush REST API 的 go 版本封装开发包,仅支持最新的REST API v3功能。
   REST API 文档：http://docs.jpush.cn/display/dev/Push-API-v3
  

使用  
----------------------------------- 
   下载源码,将jpush放入gopath下， improt "jpush/api/push".
   
   
推送流程  
----------------------------------- 
### 1.构建要推送的平台： push.PlatForm
      var pf push.PlatForm
      //pf.Add(push.ANDROID)
      pf.All
      
### 2.构建接收听众： push.Audience
      var ad push.Audience
      //s := []string{"1", "2", "3"}
      //ad.SetID(s)
      ad.All()
      
### 3.构建通知 push.AndroidNotice，或者消息： push.Message
      
      //builder : push.AndroidNotice
      var notice push.AndroidNotice
      notice.SetAlert("alert_test")
      notice.SetTitle("title_test")
      
      //或者push.Message
      var msg push.Message
      msg.Title   = "Hello"
      msg.Content = "祝大家工作顺利
      
### 4.构建builder: push.NoticeBuilder 或push.MessageBuilder 或 push.MessageAndNoticeBuilder
      //NoticeBuilder
      nb := push.NewNoticeBuilder()
      nb.SetPlatForm(&pf)
      nb.SetAudience(&ad)
      //nb.SetSimpleNotice("简单通知") //这个是简单通知，
      nb.SetAndroidNotice(&notice)
      
      //or MessageBuilder
      mb := push.NewMessageBuilder()
      mb.SetPlatForm(&pf)
      mb.SetAudience(&ad)
      mb.SetMessage(&msg)
      
### 5.构建PushClient， 发出推送
      c := push.NewPushClient(secret, appKey)
      str, err := c.Send(nb)
      //str, err := c.Send(mb)

  
### 6.未完成
      未完成ios和winphone 通知，以及不支持extras扩展字段。

