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
	pf.AddAndroid().AddQuickApp().AddIOS().AddWinPhone()
	//pf.All()

	//Audience
	var ad Audience
	s := []string{"1", "2", "3"}
	ad.SetTag(s).SetAlias(s).SetTagAnd(s).SetID([]string{registrationId}).SetTagNot(s)
	//ad.All()
	log.Println(ad.ToJson())

	//Notice
	var notice Notice
	notice.SetAlert("alert_test").
		SetAndroidNotice(&AndroidNotice{Alert: "AndroidNotice"}).
		SetIOSNotice(&IOSNotice{Alert: "IOSNotice"}).
		SetQuickAppNotice(&QuickAppNotice{Alert: "QuickAppNotice", Title: "test", Page: "url"}).
		SetWinPhoneNotice(&WinPhoneNotice{Alert: "WinPhoneNotice"})
	log.Println(notice.ToJson())

	var msg Message
	msg.SetTitle("Hello").SetMsgContent("test")
	log.Println(msg.ToJson())

	var op Option
	op.SetTimeToLive(60).SetApnsId("jiguang_test_201706011100")
	log.Println(op.ToJson())

	req := NewPushRequest()
	req.SetPlatform(&pf).
		SetAudience(&ad).
		SetMessage(&msg).
		SetNotice(&notice).
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
	result, err := client.ReportReceived([]string{msgId})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientReportStatusMessage(t *testing.T) {
	msgId := 1345223734
	result, err := client.ReportStatusMessage(&ReportStatusRequest{
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
	result, err := client.DeviceView(registrationId)
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
	result, err := client.DeviceGetWithAlias("xialei", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceDeleteAlias(t *testing.T) {
	result, err := client.DeviceDeleteAlias("xialei1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceGetTags(t *testing.T) {
	result, err := client.DeviceGetTags()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceCheckDeviceWithTag(t *testing.T) {
	result, err := client.DeviceCheckDeviceWithTag("xialei", registrationId)
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
	result, err := client.DeviceBindTags("xialei", req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceDeleteTag(t *testing.T) {
	result, err := client.DeviceDeleteTag("xialei", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientCreateScheduleTask(t *testing.T) {
	req := NewSchedule("test", "cid", true, getMsg())
	req.SingleTrigger(time.Now())
	result, err := client.ScheduleCreateTask(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
	//18785f08-c03b-11e7-be12-f8fa30f97302
}
func TestClientScheduleGetList(t *testing.T) {
	result, err := client.ScheduleGetList(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientScheduleView(t *testing.T) {
	result, err := client.ScheduleView("18785f08-c03b-11e7-be12-f8fa30f97302")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientScheduleUpdate(t *testing.T) {
	req := NewSchedule("test", "cid", false, getMsg())
	req.SingleTrigger(time.Now())

	result, err := client.ScheduleUpdate("18785f08-c03b-11e7-be12-f8fa30f97302", req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientScheduleDelete(t *testing.T) {
	result, err := client.ScheduleDelete("18785f08-c03b-11e7-be12-f8fa30f97302")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
func TestClientDevicesStatus(t *testing.T) {
	result, err := client.DeviceGetUserStatus([]string{"1104a8979215a3f999a"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
