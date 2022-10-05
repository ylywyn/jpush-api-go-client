package jpushclient

import (
	"log"
	"os"
	"testing"
	"time"
)

var client = NewClient(os.Getenv("APP_KEY"), os.Getenv("MASTER_SECRET"))
var registrationId = os.Getenv("REGISTRATION_ID")

func getMsg() *PushRequest {
	params := make(map[string]interface{})
	params["url"] = "https://www.jpush.cn"
	var pf Platform
	//pf.AddAndroid().AddQuickApp().AddIOS().AddWinPhone()
	pf.All()

	//Audience
	var ad Audience
	//s := []string{"1", "2", "3"}
	//ad.SetTag(s).SetAlias(s).SetTagAnd(s).SetID([]string{registrationId}).SetTagNot(s)
	ad.SetID([]string{registrationId})
	ad.All()
	log.Println(ad.ToJson())

	//Notice
	//mediaId:="jgmedia-1-f50da7d3-e2d5-4f34-af82-ebcb4045b4e4"
	var notice Notice
	notice.SetAlert("alert_test").
		SetAndroidNotice(&AndroidNotice{Alert: "AndroidNotice1111111", Title: "testi1111"}).
		SetIOSNotice(&IOSNotice{Alert: "IOSNotice"}).
		SetQuickAppNotice(&QuickAppNotice{Alert: "QuickAppNotice", Title: "test", Page: "/page"}).
		SetWinPhoneNotice(&WinPhoneNotice{Alert: "WinPhoneNotice"})
	log.Println(notice.ToJson())

	var notice3rd Notice3rd
	notice3rd.SetTitle("alert_testiii").SetContent("AndroidNotice1111")
	var msg Message
	msg.SetTitle("Hello").SetMsgContent("test")
	log.Println(msg.ToJson())

	var op Option
	op.SetTimeToLive(86400).SetApnsId("jiguang_test_201706011100")
	log.Println(op.ToJson())
	var cal CallBack
	cal.SetUrl("http://penguin.test.analytics.zjweishi.com").SetType("3")
	req := NewPushRequest()
	req.SetPlatform(&pf).
		SetAudience(&ad).
		SetMessage(&msg).
		SetInAppMessage(true).
		SetNotice(&notice).
		//SetCallback(&cal).
		//SetNotice3rd(&notice3rd).
		SetOptions(&op)
	log.Println(req.ToJson())
	return req
}

func TestPushRequestExecuteWithClient(t *testing.T) {
	t.Logf("test with APP_KEY: %s, MASTER_SECRET: %s", client.AppKey, client.MasterSecret)
	req := getMsg()
	result, err := client.Push(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientGetCidPool(t *testing.T) {
	data, err := client.GetCidPool(0, "push")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestClientDoPushGroup(t *testing.T) {
	result, err := client.GroupPush(getMsg())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientValidate(t *testing.T) {
	result, err := client.Validate(getMsg())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientReportReceived(t *testing.T) {
	msgId := "1345223734"
	result, err := client.GetReceived([]string{msgId})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientReportStatusMessage(t *testing.T) {
	msgId := 1345223734
	result, err := client.GetMessageStatus(&ReportStatusRequest{
		MsgId:           msgId,
		RegistrationIds: []string{registrationId},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceView(t *testing.T) {
	result, err := client.GetDevices(registrationId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceRequest(t *testing.T) {
	tags := &DeviceSettingRequestTags{
		Add: []string{"test"},
	}
	req := &DeviceSettingRequest{
		Alias:  "xialei",
		Mobile: "13333333333",
		Tags:   tags,
	}
	result, err := client.DeviceRequest(registrationId, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(result))
}

func TestClientDeviceEmptyTagsRequest(t *testing.T) {
	req := &DeviceSettingEmptyTagsRequest{
		Alias:  "xialei",
		Mobile: "13333333333",
		Tags:   "",
	}
	result, err := client.DeviceEmptyTagsRequest(registrationId, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(result))
}

func TestClientDeviceGetWithAlias(t *testing.T) {
	result, err := client.GetAliasDevices("xialei", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceDeleteAlias(t *testing.T) {
	result, err := client.DeleteAlias("xialei1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceGetTags(t *testing.T) {
	result, err := client.GetTags()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceCheckDeviceWithTag(t *testing.T) {
	result, err := client.IsDeviceInTag("xialei", registrationId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
func TestClientDeviceBindTags(t *testing.T) {
	req := &DeviceBindTagsRequest{
		Add: []string{registrationId},
	}
	result, err := client.UpdateTag("xialei", req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceDeleteTag(t *testing.T) {
	result, err := client.DeleteTag("xialei", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientCreateScheduleTask(t *testing.T) {
	req := &ScheduleRequest{
		Name:    "test",
		Enabled: true,
		Trigger: &ScheduleTrigger{
			Single: &ScheduleTriggerSingle{
				Timer: "2017-11-04 10:00:00",
			},
		},
		Push: getMsg(),
	}
	result, err := client.CreateSingleSchedule(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
	//18785f08-c03b-11e7-be12-f8fa30f97302
}
func TestClientScheduleGetList(t *testing.T) {
	result, err := client.GetSchedules(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientScheduleView(t *testing.T) {
	result, err := client.GetSchedule("18785f08-c03b-11e7-be12-f8fa30f97302")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientScheduleUpdate(t *testing.T) {
	req := &ScheduleRequest{
		Name:    "test",
		Enabled: false,
		Trigger: &ScheduleTrigger{
			Single: &ScheduleTriggerSingle{
				Timer: time.Now().String(),
			},
		},
		Push: getMsg(),
	}

	result, err := client.UpdateSingleSchedule("18785f08-c03b-11e7-be12-f8fa30f97302", req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientScheduleDelete(t *testing.T) {
	result, err := client.DeleteSchedule("18785f08-c03b-11e7-be12-f8fa30f97302")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
func TestClientDevicesStatus(t *testing.T) {
	result, err := client.GetDevicesStatus([]string{"1104a8979215a3f999a"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
func TestClientAddImageUrl(t *testing.T) {
	req := &ImageRequest{
		ImageType: 1,
		ImageUrl:  "https://pic1.zhimg.com/v2-b057cf8dfb12732399c42f658a862c34_r.jpg",
	}
	result, err := client.AddImageUrl(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
